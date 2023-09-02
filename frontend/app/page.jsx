"use client"
import * as React from 'react';
import { HandleLoginUser } from "./api/authHandler";
import Link from 'next/link';

export default function Home() {
  const [values, setValues] = React.useState({
    username: "",
    password: "",
  })

  async function handleSubmit(e) {
      e.preventDefault();
      const {username, password} = values;
      try {
        HandleLoginUser(username, password);
      } catch (err) {
        console.error(err);
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
            <h2 className="text-4xl font-mono">Sign in</h2>
            <div className="border w-full "></div>
          </div>
          <div className="m-0 p-0 h-auto w-7/12 mt-2">
            <form onSubmit={handleSubmit} className="flex flex-col w-auto">
              <input 
                type="text" 
                placeholder="username" 
                className="mt-10 bg-transparent border border-teal-400 rounded-lg p-2 font-thin"
                name="username"
                onChange={handleChange}
              />
              <input 
                type="password" 
                placeholder="********" 
                className="mt-5 bg-transparent border border-teal-400 rounded-lg p-2 font-thin"
                name="password"
                onChange={handleChange}
              />
              <div className="flex items-center justify-center mt-4">
                <button type="submit" className="border-2 border-teal-700 bg-teal-700 w-1/4 p-1 rounded-md hover:bg-teal-800 hover:border-teal-800"><Link href="/dashboard">Submit</Link></button>
              </div>
            </form>
            <div className="flex justify-center items-center mt-4">
              <p className="text-xs">you don't have account? <a href="/register" className="text-teal-900 hover:text-teal-800">sign up</a></p>
            </div>
          </div>
        </div>
      </div>
    </main>
  )
}
