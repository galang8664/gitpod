// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package log

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const (
	// OwnerField is the log field name of a workspace owner
	OwnerField = "userId"
	// WorkspaceField is the log field name of a workspace ID (not instance ID)
	WorkspaceField = "workspaceId"
	// InstanceField is the log field name of a workspace instance ID
	InstanceField = "instanceId"
	// ProjectField is the log field name of the project
	ProjectField = "projectId"
	// TeamField is the log field name of the team
	TeamField = "teamId"
)

// OWI builds a structure meant for logrus which contains the owner, workspace and instance.
// Beware that this refers to the terminology outside of wsman which maps like:
//
//	owner = owner, workspace = metaID, instance = workspaceID
func OWI(owner, workspace, instance string) log.Fields {
	return log.Fields{
		OwnerField:     owner,
		WorkspaceField: workspace,
		InstanceField:  instance,
	}
}

// LogContext builds a structure meant for logrus which contains the owner, workspace and instance.
// Beware that this refers to the terminology outside of wsman which maps like:
//
//	owner = owner, workspace = metaID, instance = workspaceID
func LogContext(owner, workspace, instance, project, team string) log.Fields {
	logFields := log.Fields{}

	if owner != "" {
		logFields[OwnerField] = owner
	}

	if workspace != "" {
		logFields[WorkspaceField] = workspace
	}

	if instance != "" {
		logFields[InstanceField] = instance
	}

	if project != "" {
		logFields[ProjectField] = project
	}

	if team != "" {
		logFields[TeamField] = team
	}

	return logFields
}

// ServiceContext is the shape required for proper error logging in the GCP context.
// See https://cloud.google.com/error-reporting/reference/rest/v1beta1/ServiceContext
// Note that we musn't set resourceType for reporting errors.
type ServiceContext struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

// Log is the application wide console logger
var Log = log.WithFields(log.Fields{})

func New() *log.Entry {
	return Log.Dup()
}

// setup default log level for components without initial invocation of log.Init.
func init() {
	logLevelFromEnv()
}

func logLevelFromEnv() {
	level := os.Getenv("LOG_LEVEL")
	if len(level) == 0 {
		level = "info"
	}

	newLevel, err := logrus.ParseLevel(level)
	if err != nil {
		Log.WithError(err).Errorf("cannot change log level to '%v'", level)
		return
	}

	Log.Logger.SetLevel(newLevel)
}

// Init initializes/configures the application-wide logger
func Init(service, version string, json, verbose bool) {
	Log = log.WithFields(log.Fields{
		"serviceContext": ServiceContext{service, version},
	})
	log.SetReportCaller(true)

	if json {
		Log.Logger.SetFormatter(&gcpFormatter{
			log.JSONFormatter{
				FieldMap: log.FieldMap{
					log.FieldKeyMsg: "message",
				},
				CallerPrettyfier: func(f *runtime.Frame) (string, string) {
					s := strings.Split(f.Function, ".")
					funcName := s[len(s)-1]
					return funcName, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
				},
				TimestampFormat: time.RFC3339Nano,
			},
		})
	} else {
		Log.Logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339Nano,
			FullTimestamp:   true,
		})
	}

	// update default log level
	logLevelFromEnv()

	if verbose {
		Log.Logger.SetLevel(log.DebugLevel)
	}
}

// gcpFormatter formats errors according to GCP rules, see
type gcpFormatter struct {
	log.JSONFormatter
}

func (f *gcpFormatter) Format(entry *log.Entry) ([]byte, error) {
	hasError := false
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			//
			// Print errors verbosely to get stack traces where available
			entry.Data[k] = fmt.Sprintf("%+v", v)
			hasError = true
		}
	}
	// map to gcp severity. See https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
	var severity string = "INFO"
	switch entry.Level {
	case logrus.TraceLevel:
		severity = "DEBUG"
	case logrus.DebugLevel:
		severity = "DEBUG"
	case logrus.InfoLevel:
		severity = "INFO"
	case logrus.WarnLevel:
		severity = "WARNING"
	case logrus.ErrorLevel:
		severity = "ERROR"
	case logrus.FatalLevel:
		severity = "CRITICAL"
	case logrus.PanicLevel:
		severity = "EMERGENCY"
	}
	entry.Data["severity"] = severity
	if entry.Level <= log.WarnLevel && hasError {
		entry.Data["@type"] = "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent"
	}

	return f.JSONFormatter.Format(entry)
}

// FromBuffer extracts the output generated by a command
// containing JSON output, parsing and writing underlying
// log data stream.
func FromBuffer(buf *bytes.Buffer, logger *logrus.Entry) {
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		b := bytes.Trim(scanner.Bytes(), "\x00")
		if len(b) == 0 {
			continue
		}

		var entry jsonEntry
		if err := json.Unmarshal(b, &entry); err != nil {
			if _, ok := err.(*json.SyntaxError); !ok {
				Log.Errorf("log.FromReader decoding JSON: %v", err)
			}

			continue
		}

		// common field name
		message := entry.Message
		if message == "" {
			// msg is defined in runc
			message = entry.Msg
		}

		// do not log empty messages
		if message == "" {
			continue
		}

		logEntry := logger.Dup()
		logEntry.Level = entry.Level
		logEntry.Message = message
		if entry.Time != nil {
			logEntry.Time = *entry.Time
		} else {
			logEntry.Time = time.Now()
		}

		// check the log of the entry is enable for the logger
		if logEntry.Logger.IsLevelEnabled(entry.Level) {
			b, err := logEntry.Bytes()
			if err != nil {
				Log.Errorf("Failed to write to custom log, %v", err)
			}
			if _, err := logEntry.Logger.Out.Write(b); err != nil {
				Log.Errorf("Failed to write to custom log, %v", err)
			}
		}
	}
}

type jsonEntry struct {
	Level   logrus.Level `json:"level,omitempty"`
	Message string       `json:"message,omitempty"`
	Msg     string       `json:"msg,omitempty"`
	Time    *time.Time   `json:"time,omitempty"`
}
