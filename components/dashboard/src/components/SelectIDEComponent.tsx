/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { IDEOption, IDEOptions } from "@gitpod/gitpod-protocol/lib/ide-protocol";
import { useContext, useEffect, useState } from "react";
import { getGitpodService } from "../service/service";
import { UserContext } from "../user-context";
import CheckBox from "./CheckBox";
import DropDown2, { DropDown2Element } from "./DropDown2";

interface SelectIDEComponentProps {
    ideOptions?: IDEOptions;
    selectedIdeOption?: string;
    useLatest?: boolean;
    setUseLatest?: (useLatestVersion: boolean) => void;
    onSelectionChange: (ide: string) => void;
}

export default function SelectIDEComponent(props: SelectIDEComponentProps) {
    const { user } = useContext(UserContext);
    function options2Elements(
        ideOptions?: IDEOptions,
        useLatest?: boolean,
        setUseLatest?: (useLatest: boolean) => void,
    ): DropDown2Element[] {
        if (!ideOptions) {
            return [];
        }
        return [
            ...IDEOptions.asArray(ideOptions).map((ide) => ({
                id: ide.id,
                element: <IdeOptionElementSelected option={ide} useLatest={!!useLatest} />,
                elementInDropDown: <IdeOptionElementInDropDown option={ide} useLatest={!!useLatest} />,
                searchableString: `${ide.label}${ide.title}${ide.notes}${ide.id}`,
                isSelectable: true,
            })),
            {
                id: "non-selectable-checkbox",
                element: (
                    <div className="ml-3">
                        <CheckBox
                            title="Use latest"
                            desc="Use latest version of IDE"
                            checked={!!useLatest}
                            onChange={(e) => setUseLatest && setUseLatest(e.target.checked)}
                        />
                    </div>
                ),
                isSelectable: false,
            },
        ];
    }

    const [elements, setElements] = useState<DropDown2Element[]>(
        options2Elements(props.ideOptions, props.useLatest, props.setUseLatest),
    );
    const [selectedIdeOption, setSelectedIdeOption] = useState<string | undefined>(
        props.selectedIdeOption || user?.additionalData?.ideSettings?.defaultIde,
    );
    useEffect(() => {
        (async () => {
            let options = props.ideOptions;
            if (!options) {
                options = await getGitpodService().server.getIDEOptions();
            }
            setElements(options2Elements(options, props.useLatest, props.setUseLatest));
            if (!selectedIdeOption || !options.options[selectedIdeOption]) {
                setSelectedIdeOption(options.defaultIde);
            }
        })();
    }, [props.ideOptions, selectedIdeOption, props.useLatest, props.setUseLatest]);
    return (
        <DropDown2
            elements={elements}
            onSelectionChange={props.onSelectionChange}
            searchable={true}
            selectedElement={selectedIdeOption}
        />
    );
}

interface IdeOptionElementProps {
    option: IDEOption | undefined;
    useLatest: boolean;
}

function IdeOptionElementSelected(p: IdeOptionElementProps): JSX.Element {
    const { option, useLatest } = p;
    if (!option) {
        return <></>;
    }
    const version = useLatest ? option.latestImageVersion : option.imageVersion;
    const label = option.type === "desktop" ? "" : option.type;

    return (
        <div className="flex" title={option.title}>
            <div className="mx-2 my-3">
                <img className="w-8 filter-grayscale self-center" src={option.logo} alt="logo" />
            </div>
            <div className="flex-col ml-1 mt-1 flex-grow">
                <div className="flex">Editor</div>
                <div className="flex text-sm text-gray-100 dark:text-gray-600">
                    <div>{option.title}</div>
                    <div className="ml-1">{version}</div>
                    <div className="ml-1">
                        {label ? (
                            <span className={`font-semibold text-sm text-gray-600 dark:text-gray-500"}`}>
                                {label[0].toLocaleUpperCase() + label.slice(1)}
                            </span>
                        ) : (
                            <></>
                        )}
                    </div>
                </div>
            </div>
        </div>
    );
}

function IdeOptionElementInDropDown(p: IdeOptionElementProps): JSX.Element {
    const { option, useLatest } = p;
    if (!option) {
        return <></>;
    }
    const version = useLatest ? option.latestImageVersion : option.imageVersion;
    const label = option.type === "desktop" ? "" : option.type;

    return (
        <div className="flex" title={option.title}>
            <div className="mx-2 my-3">
                <img className="w-8 filter-grayscale self-center" src={option.logo} alt="logo" />
            </div>
            <div className="flex-col ml-1 mt-1 flex-grow">
                <div className="flex">
                    <div>{option.title}</div>
                    <div className="ml-1 text-gray-100 dark:text-gray-600">{version}</div>
                </div>
                <div className="">
                    {label ? (
                        <span className={`font-semibold text-sm text-gray-600 dark:text-gray-500"}`}>
                            {label[0].toLocaleUpperCase() + label.slice(1)}
                        </span>
                    ) : (
                        <></>
                    )}
                </div>
            </div>
        </div>
    );
}
