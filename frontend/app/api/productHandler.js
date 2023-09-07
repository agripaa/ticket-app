import api from "./api";

const PRODUCT_ROUTES = {
    GET_ALL_PRODUCT: '/products',
    PRODUCT_ROUTE: '/product',
}

async function getProducts(){
    try {
        const dataProducts = await api.get(PRODUCT_ROUTES.GET_ALL_PRODUCT, {withCredentials: true, next: {revalidate: 60}});
        return dataProducts;
    } catch (error) {
        console.error(error);
    }
}

async function getProductById(id){
    try {
        const product = await api.get(`/${id}/${PRODUCT_ROUTES.PRODUCT_ROUTE}`, {withCredentials: true})
        return product;
    } catch (error) {
        console.error(error);
    }
}

async function createProduct(values) {
    const {name_product, desc, price} = values;

    const formData = new FormData();

    formData.append("name_product", name_product);
    formData.append("desc", desc);
    formData.append("price", price);
    try {
        const createProduct = await api.post(PRODUCT_ROUTES.PRODUCT_ROUTE, formData);
        console.log(createProduct)
        return createProduct;
    } catch (error) {
        console.error(error);
    }
}

export {getProducts, getProductById, createProduct}