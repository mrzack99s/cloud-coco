import { Api, ApiConfig } from "../api"
import { useContext, createContext } from "react";

type ContextProps = {
    instance: Api<unknown>
}

export const ApiContext = createContext<ContextProps | null>(null);

interface ApiPros {
    options: ApiConfig
    children: any
}

const ApiProvider: React.FC<ApiPros> = ({ options, children }) => {

    const apiConnector = new Api(options)

    return <ApiContext.Provider value={{ instance: apiConnector }}>{children}</ApiContext.Provider>;
}

export default ApiProvider;
export const useApiConnector = () => {
    const apiContext = useContext(ApiContext) as ContextProps
    return [ apiContext.instance ];
}