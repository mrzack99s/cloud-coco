import { FC } from "react";

interface props { }

const Overview: FC<props> = ({ }) => {

    return (
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
                                <p className="-mt-1">Create a resource</p>
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
                        <div>
                            <span className="block text-500 font-medium mb-3">
                                Resource Pools
                            </span>
                            <div className="text-900 font-medium text-xl">152</div>
                        </div>
                        <div
                            className="flex align-items-center justify-content-center bg-blue-100 border-round"
                            style={{ width: "2.5rem", height: "2.5rem" }}
                        >
                            <i className="pi pi-server text-blue-500 text-xl"></i>
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
                        <div>
                            <span className="block text-500 font-medium mb-3">
                                Resources
                            </span>
                            <div className="text-900 font-medium text-xl">152</div>
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
            <div className="col-12 md:col-6 lg:col-3">
                <div
                    className="surface-0 p-3 border-1 border-50 border-round"
                    style={{ height: "110px" }}
                >
                    <div className="flex justify-content-between mb-3">
                        <div>
                            <span className="block text-500 font-medium mb-3">
                                Resources Active
                            </span>
                            <div className="text-900 font-medium text-xl">1 / 1</div>
                        </div>
                        <div
                            className="flex align-items-center justify-content-center bg-blue-100 border-round"
                            style={{ width: "2.5rem", height: "2.5rem" }}
                        >
                            <i className="pi pi-compass text-blue-500 text-xl"></i>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Overview;