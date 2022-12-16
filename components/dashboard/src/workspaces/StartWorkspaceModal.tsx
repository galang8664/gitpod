/**
 * Copyright (c) 2021 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { useContext, useEffect, useState } from "react";
import { useLocation } from "react-router";
import Modal from "../components/Modal";
import RepositoryFinder from "../components/RepositoryFinder";
import SelectIDEComponent from "../components/SelectIDEComponent";
import { UserContext } from "../user-context";
import { StartWorkspaceModalContext } from "./start-workspace-modal-context";

export function StartWorkspaceModal() {
    const { user } = useContext(UserContext);
    const { isStartWorkspaceModalVisible, setIsStartWorkspaceModalVisible } = useContext(StartWorkspaceModalContext);
    const location = useLocation();
    const [useLatestIde, setUseLatestIde] = useState(!!user?.additionalData?.ideSettings?.useLatestVersion);

    // Close the modal on navigation events.
    useEffect(() => {
        setIsStartWorkspaceModalVisible(false);
    }, [location]);

    return (
        // TODO: Use title and buttons props
        <Modal
            onClose={() => setIsStartWorkspaceModalVisible(false)}
            onEnter={() => false}
            visible={!!isStartWorkspaceModalVisible}
        >
            <h3 className="pb-2">New Workspace</h3>
            <div className="border-t border-gray-200 dark:border-gray-800 mt-2 -mx-6 px-6 pt-4">
                <RepositoryFinder />
            </div>

            <div className="border-t border-gray-200 dark:border-gray-800 mt-2 -mx-6 px-6 pt-4">
                <div className="text-sm text-gray-100 dark:text-gray-600">Configure Workspace</div>
                <SelectIDEComponent
                    onSelectionChange={(e) => {
                        console.log(e);
                    }}
                    useLatest={useLatestIde}
                    setUseLatest={setUseLatestIde}
                />
            </div>
        </Modal>
    );
}
