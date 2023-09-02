import api from "./api";

const AUTH_ROUTES = {
    LOGIN: '/login',
    REGISTER: '/register',
    LOGOUT: '/logout',
}

async function HandleLoginUser(username, password){
    const formdata = new FormData();
    formdata.append('username', username);
    formdata.append('password', password);
    
    try {
        const loginData = await api.post(AUTH_ROUTES.LOGIN, formdata)
        window.location.href('/dashboard')
        return loginData;
    } catch (err) {
        console.error(err);
    }
}

async function HandleRegisterUser(name, username, password){
    const formdata = new FormData();
    formdata.append('name', name);
    formdata.append('username;', username);
    formdata.append('password', password);
    
    try {
        const registData = await api.post(AUTH_ROUTES.REGISTER, formdata)
        return registData;
    } catch (err) {
        console.error(err);
    }
}

async function HandleLogoutUser(){
    try {
        await api.delete(AUTH_ROUTES.LOGOUT, {withCredentials:true})
        .then(({data}) =>{
            console.log(data);
            alert('logout successfully')
        }).catch((err) =>{
            console.error(err);
        })
    } catch (err) {
        console.error(err);
    }
}

export {HandleLoginUser, HandleRegisterUser, HandleLogoutUser}