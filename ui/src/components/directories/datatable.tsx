import React, { useState, useEffect, useRef } from "react";
import { DataTable, DataTablePFSEvent } from "primereact/datatable";
import { Column } from "primereact/column";
import { useApiConnector } from "../../utils/api_connector";
import { Dropdown, DropdownChangeParams } from 'primereact/dropdown';
import { InputText } from 'primereact/inputtext';
import { Chip } from 'primereact/chip';
import { ContextMenu } from 'primereact/contextmenu'
import { Button } from 'primereact/button'

const DataTableLazy = () => {
    const [loading, setLoading] = useState(false);
    const [totalRecords, setTotalRecords] = useState(0);
    const [data, setData] = useState([] as any);
    const [pagiOption, setPagiOption] = useState({
        first: 0,
        perPage: 10,
    });
    const [apiInstance] = useApiConnector()
    const [selectedCtxRow, setSelectedCtxRow] = useState({});

    const menuModel = [
        {
            label: 'View', icon: 'pi pi-pencil', command: () => {
                console.log(selectedCtxRow)
            }
        },
        { label: 'Delete', icon: 'pi pi-trash' }
    ];
    const cm = useRef({} as any);

    useEffect(() => {
        loadLazyData();
    }, [pagiOption]); // eslint-disable-line react-hooks/exhaustive-deps

    const loadLazyData = () => {
        setLoading(true);

        apiInstance.directories.getDirectoriesByOffset(
            {
                offset: pagiOption.first + 1,
                limit: pagiOption.perPage
            }
        )
            .then(response => response.data)
            .then((data) => {
                setTotalRecords(data.record_count!);
                setData(data.record_list);
                setLoading(false);
            })
    };

    const template2 = {
        layout: 'RowsPerPageDropdown CurrentPageReport PrevPageLink NextPageLink',
        'RowsPerPageDropdown': (options: { value: any; onChange: ((e: DropdownChangeParams) => void) | undefined; }) => {
            const dropdownOptions = [
                { label: 10, value: 10 },
                { label: 20, value: 20 },
                { label: 50, value: 50 }
            ];

            return (
                <React.Fragment>
                    <Dropdown value={options.value} options={dropdownOptions} onChange={options.onChange} />
                </React.Fragment>
            );
        },
        'CurrentPageReport': (options: { first: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | React.ReactFragment | React.ReactPortal | null | undefined; last: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | React.ReactFragment | React.ReactPortal | null | undefined; totalRecords: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | React.ReactFragment | React.ReactPortal | null | undefined; }) => {
            return (
                <span style={{ color: 'var(--text-color)', userSelect: 'none', width: '120px', textAlign: 'center' }}>
                    {options.first} - {options.last} of {options.totalRecords}
                </span>
            )
        }
    };

    const onPage = (e: DataTablePFSEvent) => {
        setPagiOption({
            first: e.first!,
            perPage: e.rows
        })
    };



    return (
        <>
            <ContextMenu model={menuModel} ref={cm} onHide={() => setSelectedCtxRow({})} />
            <div className="grid w-full">
                <div className="col-12 text-color-secondary">
                    <span className="flex justify-content-start font-bold">
                        Directories
                    </span>
                    <span className="flex justify-content-end -mt-3">
                        <span className="p-input-icon-right">
                            <i className="pi pi-search" />
                            <InputText className="p-inputtext-sm" placeholder="Search" />
                        </span>
                    </span>
                </div>
                <div className="col-12">
                    <div className="card">
                        <DataTable
                            value={data}
                            lazy
                            onContextMenu={e => cm.current.show(e.originalEvent)}
                            contextMenuSelection={selectedCtxRow}
                            onContextMenuSelectionChange={e => setSelectedCtxRow(e.value)}
                            responsiveLayout="scroll"
                            dataKey="id"
                            paginator
                            paginatorTemplate={template2.layout}
                            first={pagiOption.first}
                            rows={pagiOption.perPage}
                            totalRecords={totalRecords}
                            stripedRows
                            onPage={onPage}
                            loading={loading}
                            rowsPerPageOptions={[10, 20, 50]}
                            paginatorClassName="justify-content-end"
                        >
                            <Column
                                field="name"
                                header="Name"
                                style={{
                                    width: "30%"
                                }}
                            />
                            <Column
                                field="uuid"
                                header="UUID"
                                style={{
                                    width: "30%"
                                }}
                            />

                            <Column
                                style={{
                                    width: "40%"
                                }}
                                className="text-right"
                                body={(e) => (
                                    <>
                                        <Button icon="pi pi-ellipsis-v" className="p-button-rounded p-button-primary p-button-text"  onClick={(ee) => {
                                            setSelectedCtxRow(e)
                                            cm.current.show(ee)
                                        }} />
                                    </>
                                )}
                            />

                        </DataTable>
                    </div>
                </div>
            </div>
        </>

    );
};

export default DataTableLazy;
