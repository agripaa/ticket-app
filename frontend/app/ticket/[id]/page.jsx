"use client"
import * as React from 'react';
import { getProductById } from "../../api/productHandler"
import Navbar from '@/app/layout/navbar';

export default function Ticket({params}) {
    const [ticket, setTicket] = React.useState({});
    const {id} = params;

    async function getProduct(){
        try {
            const ticket = await getProductById(id);
            const ticketData = await ticket.data;
            setTicket(ticketData);
        } catch (err) {
            console.error(err);
        }
    }

    React.useEffect(() => {
        getProduct();
    }, []);

    return (
        <Navbar>
            <section className="w-full h-[80vh] px-8 mt-6">
                <div className="flex flex-wrap justify-center">
                    <div className="flex flex-col items-center m-2 bg-white border border-gray-200 rounded-lg shadow md:flex-row md:max-w-xl hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700">
                        <div className="flex flex-col justify-between p-4 leading-normal">
                            <h5 className="mb-2 text-center text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{ticket.name_product}</h5>
                            <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">{ticket.desc}</p>
                            <p className="text-end mr-3 font-light text-gray-400 dark:text-orange-200 ">{ticket.price}$</p>
                        </div>
                    </div>
                </div>
            </section>
        </Navbar>
    )
}
