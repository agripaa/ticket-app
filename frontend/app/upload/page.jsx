"use client"

import * as React from "react"
import { createProduct } from "../api/productHandler";
import Link from "next/link";
import Navbar from "../layout/navbar";

export default function upload () {
    const [values, setValues] = React.useState({
      name_product: "",
      desc: "",
      price: 0,
    })

    async function handleSubmit(e){
      e.preventDefault();
      try {
        createProduct(values);
        alert("created data successfully!");
      } catch (err) {
        console.error(err);
      }
    }

    function handleChange(e) {
      setValues({...values, [e.target.name]: e.target.value})
    }
    return (
        <>
          <Navbar>
            
            <main className="flex min-h-screen items-center justify-center">
              <div className="flex justify-center items-center w-4/12">
                <div className="flex flex-col border-2 border-teal-600 rounded-md pt-20 pb-16 w-full h-3/5 items-center justify-center">
                  <div>
                    <h2 className="text-4xl font-mono">Upload Ticket</h2>
                    <div className="border w-full "></div>
                  </div>
                  <div className="m-0 p-0 h-auto w-7/12 mt-2">
                    <form onSubmit={handleSubmit} className="flex flex-col w-auto">
                      <input 
                        type="text" 
                        placeholder="Name Product" 
                        className="mt-10 bg-transparent border border-teal-400 rounded-lg p-2 font-thin"
                        name="name_product"
                        onChange={handleChange}
                      />
                      <input 
                        type="number" 
                        placeholder="$ xxx" 
                        className="mt-5 bg-transparent border border-teal-400 rounded-lg p-2 font-thin"
                        name="price"
                        onChange={handleChange}
                        min='1'
                        max='1000000000'
                      />
                      <textarea 
                        type="text" 
                        placeholder="Description"  
                        className="mt-5 bg-transparent border border-teal-400 rounded-lg p-2 font-thin"
                        name="desc"
                        onChange={handleChange}
                      />
                      <div className="flex items-center justify-center mt-4">
                          <button type="submit" className="border-2 border-teal-700 bg-teal-700 w-1/4 p-1 rounded-md hover:bg-teal-800 hover:border-teal-800">Submit</button>
                      </div>
                    </form>
                  </div>
                </div>
              </div>
            </main>
          </Navbar>
        </>
    )
}