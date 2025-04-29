import { createContext, ReactNode, useCallback, useEffect, useState } from "react";
import axios from "axios";

interface AuthContextType {
    loggedIn: boolean;
    loadingUser: boolean;
    username?: string;
    logout: () => Promise<void>;
    checkAuthentication: () => Promise<void>;
    availableProviders: ApiTypes.InfoProviderTypes[];
};

const AppContext = createContext<AuthContextType>({} as AuthContextType);

const AppContextProvider = (props: { children?: ReactNode | undefined }) => {
    const [username, setUsername] = useState("");
    const [loadingUser, setLoadingUser] = useState(true);
    const [loggedIn, setLoggedIn] = useState(false);
    const [availableProviders, setAvailableProviders] = useState<ApiTypes.InfoProviderTypes[]>([]);

    const logout = useCallback(async () => {
        try {
            await axios.post("/api/account/logout");
            setLoggedIn(false);
            setUsername("");
            setLoadingUser(false);
        } catch (error) {
            console.error("Logout failed:", error);
        }
    }, [setLoggedIn, setUsername]);

    const checkAuthentication = useCallback(async () => {
        try {
            const statusData = await axios.get<ApiTypes.AccountUserResponse>("/api/account/user");
            if (!statusData) {
                // the user is not logged in
                setLoggedIn(false);
                setUsername("");
                return;
            }
            setLoggedIn(true);
            setUsername(statusData.data.username);
        } catch (error) {
            console.error("Authentication check failed:", error);
            setLoggedIn(false);
            setUsername("");
            setLoadingUser(false);
        }
    }, [setUsername, setLoggedIn]);

    const fetchServerInfo = useCallback(async () => {
        const info = await axios.get<ApiTypes.InfoResponse>("/api/info");

        setAvailableProviders(info.data.supported_oauth_providers);
    }, []);

    useEffect(() => {
        setLoadingUser(true);
        fetchServerInfo().then();
        checkAuthentication().then();
    }, [checkAuthentication, fetchServerInfo]);

    window.addEventListener("storage", () => {
        checkAuthentication().then();
    });

    return (
        <AppContext.Provider
            value={{ loggedIn, loadingUser, username, logout, checkAuthentication, availableProviders }}
        >
            {props.children}
        </AppContext.Provider>
    );
};

export { AppContext, AppContextProvider };
