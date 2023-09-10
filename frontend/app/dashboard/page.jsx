"use client"

import * as React from "react";
import { getProducts } from '../api/productHandler.js';
import { useRouter } from "next/navigation";
import Navbar from "../layout/navbar.jsx";

export default function Dashboard (){
    const [ticketProduct, setTicketProduct] = React.useState([]);
    const { push } = useRouter();

    async function getDataTicket () {
        try {
            const ticketDatas = await getProducts();
            const tickets = ticketDatas.data;

            setTicketProduct(tickets);
        } catch (error) {
            console.error("error : ", error);
        }
    }
    React.useEffect(() => {
        getDataTicket();
    }, [])
    return (
        <>
            <Navbar>
                <section className="w-full h-full px-8 mt-6">
                    <div className="flex flex-wrap justify-center">
                        {ticketProduct.status == '401' ? push('/') : ticketProduct.map((ticket, i) => {
                            return (
                                <>
                                    <a href={`/ticket/${ticket.id}`} class="flex flex-col items-center m-2 bg-white border border-gray-200 rounded-lg shadow md:flex-row md:max-w-xl hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700" key={i}>
                                        <div class="flex flex-col justify-between p-4 leading-normal">
                                            <h5 class="mb-2 text-center text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{ticket.name_product}</h5>
                                            <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">{ticket.desc}</p>
                                            <p class="text-end mr-3 font-light text-gray-400 dark:text-orange-200 ">{ticket.price}$</p>
                                        </div>
                                    </a>
                                </>
                            )
                        })}
                    </div>
                </section>
            </Navbar>
        </>
    )
}