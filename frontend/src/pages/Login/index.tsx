import React, { useState, useEffect } from "react";
import { Navigate, useNavigate } from "react-router-dom";
import axios from "axios";
import { ExclamationTriangleIcon } from "@heroicons/react/24/solid";
//import { SiCodeberg, SiDiscord, SiGithub } from "react-icons/si";
//import { AppContext } from "../../services/AppContext.tsx";
import ApiUrls from "../../services/api_urls";
// import cdbLogo from "../../assets/content_db_logo.png";

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const [canSubmit, setCanSubmit] = useState(false);

    const [loading, setLoading] = useState(false);
    const [err, setErr] = useState<string | null>(null);

    //const { loggedIn, availableProviders } = useContext(AppContext);

    useEffect(() => {
        if (username === "" || password === "") {
            setCanSubmit(false);
        } else {
            setCanSubmit(true);
        }
    }, [username, password, setCanSubmit]);

    const navigate = useNavigate();

    // Navigate to homepage if already authenticated
    if (true) {
        return <Navigate to="/" />;
    }

    async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        setLoading(true);
        await axios
            .post(ApiUrls.AccountLogin, {
                username: username,
                password: password,
            })
            .then(() => {
                //console.log(r);
                setLoading(false);
                navigate("/");
            })
            .catch((e: { message: string }) => {
                //console.log(e);
                setLoading(false);
                setErr(e.message);
            });
    }

    return (
        <div className="flex justify-center align-middle">
            <div className="app-pannel mx-0 my-10 flex w-full max-w-md flex-col p-6 md:mx-10">
                <h1 className="panel-text-heading">Login</h1>
                <form className="mt-4 flex flex-col" onSubmit={handleSubmit}>
                    <label htmlFor="username" className="mb-2 block select-none text-sm font-medium dark:text-white">
                        Username
                    </label>
                    <input
                        type="text"
                        name="username"
                        id="username"
                        placeholder=""
                        spellCheck={false}
                        className="form-input2"
                        value={username}
                        onChange={(e) => {
                            setUsername(e.target.value);
                        }}
                    />
                    <label
                        htmlFor="password"
                        className="mb-2 mt-4 block select-none text-sm font-medium dark:text-white"
                    >
                        Password
                    </label>
                    <input
                        type="password"
                        name="password"
                        id="password"
                        placeholder=""
                        className="form-input2"
                        value={password}
                        onChange={(e) => {
                            setPassword(e.target.value);
                        }}
                    />
                    {err !== null ? (
                        <div className="mt-4 flex flex-row items-center gap-2 text-red-600">
                            <ExclamationTriangleIcon className="h-8 w-8" />
                            {err}
                        </div>
                    ) : undefined}
                    <div className="mt-4 flex gap-4">
                        <button
                            className="button-secondary w-full justify-center"
                            onClick={() => {
                                navigate("/register");
                            }}
                        >
                            Register
                        </button>
                        <button
                            className={`button-primary w-full justify-center ${loading ? "cursor-wait" : null}`}
                            type="submit"
                            disabled={!canSubmit}
                        >
                            Login
                        </button>
                    </div>
                    <hr className="my-4" />
                    {/* <div className="flex flex-col gap-4">
                        {availableProviders.includes("contentdb") ? (
                            <a
                                className="button-secondary inline-flex w-full flex-row items-center justify-start gap-x-2"
                                href={ApiUrls.AccountProviderCDB}
                            >
                                <img className={"size-5"} src={cdbLogo} alt="ContentDB Logo" />
                                <span>ContentDB</span>
                            </a>
                        ) : null}
                        {availableProviders.includes("github") ? (
                            <a
                                className="button-secondary inline-flex w-full flex-row items-center justify-start gap-x-2"
                                href={ApiUrls.AccountProviderGitHub}
                            >
                                <SiGithub className="size-5" color={"#181717"} />
                                <span>GitHub</span>
                            </a>
                        ) : null}
                        {availableProviders.includes("codeberg") ? (
                            <a
                                className="button-secondary inline-flex w-full flex-row items-center justify-start gap-x-2"
                                href={ApiUrls.AccountProviderCodeberg}
                            >
                                <SiCodeberg className="size-5" color={"#2185D0"} />
                                <span>Codeberg</span>
                            </a>
                        ) : null}
                        {availableProviders.includes("discord") ? (
                            <a
                                className="button-secondary inline-flex w-full flex-row items-center justify-start gap-x-2"
                                href={ApiUrls.AccountProviderDiscord}
                            >
                                <SiDiscord className="size-5" color={"#5865F2"} />
                                <span>Discord</span>
                            </a>
                        ) : null}
                    </div> */}
                </form>
            </div>
        </div>
    );
}

export default Login;
