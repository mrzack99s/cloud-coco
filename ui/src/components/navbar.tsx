import React, { useState } from "react";
import { Menubar } from "primereact/menubar";
import { Button } from "primereact/button";
import { Sidebar } from "primereact/sidebar";
import { Avatar } from "primereact/avatar";
import { Divider } from "primereact/divider";
import './custom.css'
import { useCookies } from 'react-cookie';
import { useNavigate } from 'react-router-dom';

const items = [
    {
        label: "Management",
        icon: "pi pi-building",
        items: [
            {
                label: "Directories",
                icon: "pi pi-sitemap",
                url: "/directories"
            },
            {
                separator: true,
            },
            {
                label: "Subscriptions",
                icon: "pi pi-key",
            },
        ],
    },
];

export default () => {
    const [cookies, setCookie, removeCookie] = useCookies(['role']);
    const navigate = useNavigate()

    const end = () => (
        <>
            <Button
                label="CP"
                className="my-1 p-0 p-button-rounded surface-ground text-color-secondary"
                aria-label="Filter"
                style={{
                    borderRadius: "100%",
                    height: "40px",
                    width: "40px"
                }}
                onClick={() => setVisibleRight(true)}
            />
        </>
    );
    const start = () => (
        <div><Button label="COCO" onClick={() => navigate("/")} className="p-button-text text-white text-2xl p-0 bg-blue-500 mr-4" /></div>
    );
    const [visibleRight, setVisibleRight] = useState(false);


    return (
        <>
            <Sidebar
                visible={visibleRight}
                position="right"
                onHide={() => setVisibleRight(false)}
            >
                <div className="text-center">
                    <Avatar label="CP" className="mr-2" size="xlarge" shape="circle" />
                    <p>Chatdanai Phakaket</p>
                </div>
                <Divider align="left" type="dashed" className="mt-4 mb-0">
                    <span className="text-sm">Context</span>
                </Divider>
                <div className="mt-0">
                    <ul className="list-none p-0 m-0">
                        <li className="flex align-items-center py-3 px-2 border-bottom-1 border-300 flex-wrap">
                            <div className="text-500 w-4 font-medium">Directory</div>
                            <div className="text-900 w-6 md:flex-order-0 flex-order-1">
                                Heat
                            </div>
                            <div className="text-900 w-2 md:flex-order-0 text-xs flex-order-1">
                                <Button label="Change" className="p-button-link text-xs p-0" />
                            </div>
                        </li>
                    </ul>
                </div>
                <Divider align="left" type="dashed" className="mt-4 mb-0">
                    <span className="text-sm">Preferences</span>
                </Divider>
                <div className="mt-4 text-center">
                    <Button label="Change Password" className="mr-2 text-xs text-white p-button-warning" />
                    <Button label="Sign Out" className="text-xs p-button-danger" icon="pi pi-sign-out" />
                </div>
            </Sidebar>
            {cookies.role == "administrator" &&
                < Menubar
                    className="py-0 px-3 bg-blue-500"
                    start={start}
                    model={items}
                    end={end}
                    style={{ width: "100%", position: "fixed", zIndex: 100 }}
                />
            }

            {cookies.role != "administrator" &&
                < Menubar
                    className="py-0 px-3 bg-blue-500"
                    start={start}
                    end={end}
                    style={{ width: "100%", position: "fixed", zIndex: 100 }}
                />
            }

        </>
    );
};
