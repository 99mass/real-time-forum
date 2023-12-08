let routes = {
    "/Home": {
        name: "Home"
    },
    "/": {
        name: "Login"
    },
    "/Login": {
        name: "Login"
    },
    
    "/Registre": {
        name: "Registre"
    },
    "/AddPost": {
        name: "AddPost"
    },
};

const addRouter=(name)=>{
    history.pushState({}, "", `/${name}`);
}
const replaceRouter= (name)=>{
    history.replaceState({}, "", `/${name}`);
}
const titlePage=(name)=>{
    document.title = name;

}
let currentPath = window.location.pathname;

export{
    routes,addRouter,replaceRouter,currentPath,titlePage
}