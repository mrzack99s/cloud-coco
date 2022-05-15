import React, { useState, useEffect, useRef, FC } from "react";
import { DataTable, DataTablePFSEvent } from "primereact/datatable";
import { Column } from "primereact/column";
import { useApiConnector } from "../../../utils/api_connector";
import { Dropdown, DropdownChangeParams } from 'primereact/dropdown';
import { InputText } from 'primereact/inputtext';
import { Chip } from 'primereact/chip';
import { ContextMenu } from 'primereact/contextmenu'
import { Button } from 'primereact/button'
import { Dialog } from "primereact/dialog";
import { View } from "./view";
import { ModelsServices, ModelsServiceVersions } from "../../../api";

interface props {
    trigger: boolean
}

const DataTableLazy: FC<props> = ({ trigger }) => {

    const [loading, setLoading] = useState(false);
    const [totalRecords, setTotalRecords] = useState(0);
    const [data, setData] = useState([] as any);
    const [pagiOption, setPagiOption] = useState({
        first: 0,
        perPage: 10,
    });

    const [loading2, setLoading2] = useState(false);
    const [totalRecords2, setTotalRecords2] = useState(0);
    const [data2, setData2] = useState([] as any);
    const [pagiOption2, setPagiOption2] = useState({
        first: 0,
        perPage: 10,
    });

    const [apiInstance] = useApiConnector()
    const [selectedCtxRow, setSelectedCtxRow] = useState({} as ModelsServices);
    const [selectedCtxRow2, setSelectedCtxRow2] = useState({} as ModelsServiceVersions);
    const [displayViewVersion, setDisplayViewVersion] = useState(false);

    const menuModel = [
        {
            label: 'View', icon: 'pi pi-pencil', command: () => {
                setDisplayViewVersion(true)
                loadLazyDataServicesVersion()
            }
        },
        { label: 'Delete', icon: 'pi pi-trash' }
    ];
    const menuModel2 = [
        {
            label: 'View', icon: 'pi pi-pencil', command: () => {
                setDisplayViewVersion(true)
            }
        },
        { label: 'Delete', icon: 'pi pi-trash' }
    ];

    const cm = useRef({} as any);
    const cm2 = useRef({} as any);

    useEffect(() => {
        loadLazyDataServices();
    }, [pagiOption, trigger]); // eslint-disable-line react-hooks/exhaustive-deps

    const loadLazyDataServices = () => {
        setLoading(true);

        apiInstance.services.getServicesByOffset(
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

    const loadLazyDataServicesVersion = () => {
        setLoading2(true);
        console.log(selectedCtxRow)
        apiInstance.serviceVersions.getServiceVersionsBySid(
            {
                sid: selectedCtxRow.id!,
                offset: pagiOption2.first + 1,
                limit: pagiOption2.perPage
            }
        )
            .then(response => response.data)
            .then((data) => {
                console.log(data)
                setTotalRecords2(data.record_count!);
                setData2(data.record_list);
                setLoading2(false);
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

    const onPage2 = (e: DataTablePFSEvent) => {
        setPagiOption2({
            first: e.first!,
            perPage: e.rows
        })
    };



    return (
        <>
            <ContextMenu model={menuModel} ref={cm} />
            <ContextMenu model={menuModel2} ref={cm2} />
            {/* <Dialog header="Create a service" visible={displayViewDialog} style={{ width: '1000px', height: '500px' }} onHide={() => {
                setDisplayViewDialog(false)
            }}>
                <View setDisplay={setDisplayViewDialog} />
            </Dialog> */}

            {displayViewVersion &&
                <div className="grid w-full">
                    <div className="col-12 text-color-secondary">
                        <span className="flex justify-content-start font-bold">
                            {selectedCtxRow.name} versions
                            <Button label="Back to services available" className="p-button-sm p-1 -mt-1 ml-3 p-button-outlined" onClick={
                                () => {
                                    setDisplayViewVersion(false)
                                }
                            } />
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
                                value={data2}
                                lazy
                                onContextMenu={e => cm2.current.show(e.originalEvent)}
                                contextMenuSelection={selectedCtxRow2}
                                onContextMenuSelectionChange={e => setSelectedCtxRow2(e.value)}
                                responsiveLayout="scroll"
                                dataKey="id"
                                paginator
                                paginatorTemplate={template2.layout}
                                first={pagiOption2.first}
                                rows={pagiOption2.perPage}
                                totalRecords={totalRecords2}
                                stripedRows
                                onPage={onPage2}
                                loading={loading2}
                                rowsPerPageOptions={[10, 20, 50]}
                                paginatorClassName="justify-content-end"
                            >
                                <Column
                                    field="version"
                                    header="Verion"
                                    style={{
                                        width: "20%"
                                    }}
                                />
                                <Column
                                    field="uuid"
                                    header="UUID"
                                    style={{
                                        width: "20%"
                                    }}
                                />
                                <Column
                                    field="image"
                                    header="Image"
                                    style={{
                                        width: "20%"
                                    }}
                                />
                                <Column
                                    style={{
                                        width: "10"
                                    }}
                                    body={(e) => (
                                        <>
                                            <Button icon="pi pi-ellipsis-v" className="p-button-rounded p-button-primary p-button-text" onClick={(ee) => {
                                                setSelectedCtxRow2(e)
                                                cm2.current.show(ee)
                                            }} />
                                        </>
                                    )}
                                />

                            </DataTable>
                        </div>
                    </div>
                </div>
            }

            {!displayViewVersion &&
                <div className="grid w-full">
                    <div className="col-12 text-color-secondary">
                        <span className="flex justify-content-start font-bold">
                            Services Available
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
                                onContextMenuSelectionChange={e => {
                                    setSelectedCtxRow(e.value)
                                }}
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
                                        width: "20%"
                                    }}
                                />
                                <Column
                                    field="uuid"
                                    header="UUID"
                                    style={{
                                        width: "20%"
                                    }}
                                />
                                <Column
                                    header="Version"
                                    body={(e) => {
                                        let allVersion: any = []
                                        if (e.versions) {
                                            e.versions.forEach((element: any) => {
                                                allVersion.push(<Chip key={element.uuid} label={element.version} className="text-xs mr-2 mb-2 custom-chip" />)
                                            });
                                        } else {
                                            allVersion.push(<Chip key={new Date().getTime()} label="Unavailable" className="text-xs mr-2 mb-2 custom-chip" />)
                                        }
                                        return (
                                            <>
                                                {allVersion}
                                            </>
                                        )
                                    }}
                                    style={{
                                        width: "50%"
                                    }}
                                />
                                <Column
                                    style={{
                                        width: "10"
                                    }}
                                    body={(e) => (
                                        <>
                                            <Button icon="pi pi-ellipsis-v" className="p-button-rounded p-button-primary p-button-text" onClick={(ee) => {
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
            }
        </>

    );
};

export default DataTableLazy;
