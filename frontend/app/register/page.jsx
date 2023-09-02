"use client"
import * as React from "react"
import { HandleRegisterUser } from "../api/authHandler";

export default function Register(){
    const [values, setValues] = React.useState({
        name: "",
        username: "",
        password: "",
    })

    
    function handleSubmit(e) {
        e.preventDefault();
        const {name, username, password} = values;
        try {
            HandleRegisterUser(name, username, password);
            window.location.href('/')
        } catch (error) {
            console.error(error);
        }
    }

    function handleChange(e) {
        setValues({...values, [e.target.name]: e.target.value})
    }
    return (
        <main className="flex min-h-screen items-center justify-center">
        <div className="flex justify-center items-center w-4/12">
            <div className="flex flex-col border-2 border-teal-600 rounded-md pt-20 pb-16 w-full h-3/5 items-center justify-center">
            <div>
                <h2 className="text-4xl font-mono">Sign up</h2>
                <div className="border w-full "></div>
            </div>
            <div className="m-0 p-0 h-auto w-7/12 mt-2">
                <form onSubmit={handleSubmit} className="flex flex-col w-auto">
                <input
                 name="name"
                 type="text"
                 placeholder="name"
                 className="mt-10 bg-transparent border border-teal-400 rounded-lg p-2 font-thin" 
                 onChange={handleChange}
                />
                <input
                 name="username"
                 type="text"
                 placeholder="username" 
                 className="mt-5 bg-transparent border border-teal-400 rounded-lg p-2 font-thin"
                 onChange={handleChange}
                />
                <input
                 name="password"
                 type="password"
                 placeholder="********" 
                 className="mt-5 bg-transparent border border-teal-400 rounded-lg p-2 font-thin"
                 onChange={handleChange}
                />
                <div className="flex items-center justify-center mt-4">
                    <button type="submit" className="border-2 border-teal-700 bg-teal-700 w-1/4 p-1 rounded-md hover:bg-teal-800 hover:border-teal-800">Submit</button>
                </div>
                </form>
                <div className="flex justify-center items-center mt-4">
                <p className="text-xs">you've an account? <a href="/" className="text-teal-900 hover:text-teal-800">sign in</a></p>
                </div>
            </div>
            </div>
        </div>
    </main>
    )
}