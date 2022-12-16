/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { useMemo, useState } from "react";
import Arrow from "./Arrow";

export interface DropDown2Element {
    id: string;
    elementInDropDown?: JSX.Element;
    element: JSX.Element;
    searchableString?: string;
    isSelectable?: boolean;
}

export interface DropDown2Props {
    elements: DropDown2Element[];
    selectedElement?: string;
    searchable?: boolean;
    onSelectionChange: (id: string) => void;
}

export default function DropDown2(props: DropDown2Props) {
    const [selectedElement, setSelectedElement] = useState<string>(props.selectedElement || props.elements[0].id);
    const [showDropDown, setShowDropDown] = useState<boolean>(false);
    const onSelected = useMemo(
        () => (elementId: string) => {
            props.onSelectionChange(elementId);
            setSelectedElement(elementId);
            setShowDropDown(false);
        },
        [props],
    );
    const doShowDropDown = useMemo(() => () => setShowDropDown(true), []);
    if (props.elements.length === 0) {
        return <></>;
    }
    return (
        <div>
            {showDropDown ? (
                <div
                    onKeyDown={(e) => {
                        if (e.key === "Escape") {
                            setShowDropDown(false);
                            e.preventDefault();
                        }
                    }}
                >
                    <DropDownMenu {...props} onSelected={onSelected} />
                </div>
            ) : (
                <div
                    className="rounded-xl hover:bg-gray-400 dark:hover:bg-gray-700 cursor-pointer flex items-center px-2"
                    onClick={doShowDropDown}
                >
                    {props.elements.find((e) => e.id === selectedElement)?.element}
                    <div className="flex-grow" />
                    <div>
                        <Arrow direction={"down"} />
                    </div>
                </div>
            )}
        </div>
    );
}

interface DropDownMenuProps extends DropDown2Props {
    onSelected: (ide: string) => void;
}

function DropDownMenu(props: DropDownMenuProps): JSX.Element {
    const { elements, selectedElement, searchable, onSelected } = props;
    const [search, setSearch] = useState<string>("");
    const filteredOptions = useMemo(
        () =>
            elements.filter(
                (o) => !o.isSelectable || (o.searchableString || "").toLowerCase().indexOf(search.toLowerCase()) !== -1,
            ),
        [search, elements],
    );
    return (
        <div className="relative flex flex-col">
            {searchable && (
                <input
                    type="text"
                    autoFocus
                    className="w-full focus"
                    placeholder="Search IDE"
                    value={search}
                    onChange={(e) => setSearch(e.target.value)}
                />
            )}
            <div className="absolute w-full top-11 bg-gray-900 max-h-72 overflow-auto rounded-xl mt-3">
                {filteredOptions.map((element) => {
                    const selected = element.id === selectedElement;

                    let selectionClasses = `hover:bg-gray-400 dark:hover:bg-gray-700 cursor-pointer`;
                    if (selected) {
                        selectionClasses = `bg-gray-300 dark:bg-gray-800`;
                    }
                    if (!element.isSelectable) {
                        selectionClasses = ``;
                    }
                    return (
                        <div
                            key={element.id}
                            className={"rounded-md " + selectionClasses}
                            onClick={() => element.isSelectable && onSelected(element.id)}
                        >
                            {element.elementInDropDown || element.element}
                        </div>
                    );
                })}
            </div>
        </div>
    );
}
