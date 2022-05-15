import { FC } from "react";
import { ProgressBar } from 'primereact/progressbar';
import DataTableLazy from "./datatable";

interface props { }

const Overview: FC<props> = ({ }) => {

    return (
        <>
            <div className="grid">
                <div className="col-12 md:col-6 lg:col-3">
                    <div
                        className="portal-pointer surface-0 p-3  primary-box p-ripple border-1 border-50 border-round"
                        style={{ height: "110px" }}
                    >
                        <div>
                            <div
                                className="text-900 font-medium text-xl text-center"
                            >
                                <span className="p-text-secondary text-xl pt-2">
                                    <i className="pi pi-plus-circle"></i>
                                    <p className="-mt-1">Create a directory</p>
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="col-12 md:col-6 lg:col-3">
                    <div
                        className="surface-0 p-3 border-1 border-50 border-round"
                        style={{ height: "110px" }}
                    >
                        <div className="flex justify-content-between mb-3">
                            <div className="w-full">
                                <span className="block text-500 font-medium mb-3">
                                    Directories
                                </span>
                                <div className="text-900 font-medium text-xl">
                                    <ProgressBar value={50} className="w-9" style={{ height: "10px" }}></ProgressBar>
                                </div>
                                <div className="text-xs text-color-secondary mt-1">
                                    <span style={{ position: "relative" }}>
                                        <span className="font-medium">Allocated</span>
                                        <span className="text-500"> 10%, 0.12 cores used / 4 cores</span>
                                    </span>
                                </div>
                            </div>
                            <div
                                className="flex align-items-center justify-content-center bg-blue-100 border-round"
                                style={{ width: "2.5rem", height: "2.5rem" }}
                            >
                                <i className="mdi mdi-cpu-64-bit text-blue-500 text-xl"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="col-12 md:col-6 lg:col-3">
                    <div
                        className="surface-0 p-3 border-1 border-50 border-round"
                        style={{ height: "110px" }}
                    >
                        <div className="flex justify-content-between mb-3">
                            <div className="w-full">
                                <span className="block text-500 font-medium mb-3">
                                    RAM (GiB)
                                </span>
                                <div className="text-900 font-medium text-xl">
                                    <ProgressBar color={10 > 1 ? "red" : "primary"} value={50} className="w-9" style={{ height: "10px" }}></ProgressBar>
                                </div>
                                <div className="text-xs text-color-secondary mt-1">
                                    <span style={{ position: "relative" }}>
                                        <span className="font-medium">Allocated</span>
                                        <span className="text-500"> 10%, 0.29 GiB used / 4.00 GiB</span>
                                    </span>
                                </div>
                            </div>
                            <div
                                className="flex align-items-center justify-content-center bg-blue-100 border-round"
                                style={{ width: "2.5rem", height: "2.5rem" }}
                            >
                                <i className="mdi mdi-memory text-blue-500 text-xl"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="col-12 md:col-6 lg:col-3">
                    <div
                        className="surface-0 p-3 border-1 border-50 border-round"
                        style={{ height: "110px" }}
                    >
                        <div className="flex justify-content-between mb-3">
                            <div className="w-full">
                                <span className="block text-500 font-medium mb-3">
                                    Resources
                                </span>
                                <div className="text-900 font-medium text-xl">
                                    <ProgressBar value={50} className="w-9" style={{ height: "10px" }}></ProgressBar>
                                </div>
                                <div className="text-xs text-color-secondary mt-1">
                                    <span style={{ position: "relative" }}>
                                        <span className="font-medium">Allocated</span>
                                        <span className="text-500"> 10%, 1 used / 20 resources</span>
                                    </span>
                                </div>
                            </div>
                            <div
                                className="flex align-items-center justify-content-center bg-blue-100 border-round"
                                style={{ width: "2.5rem", height: "2.5rem" }}
                            >
                                <i className="pi pi-box text-blue-500 text-xl"></i>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div className="mt-3">
                <DataTableLazy />
            </div>
        </>
    )
}

export default Overview;