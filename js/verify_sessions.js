import { linkApi } from "./helper/api_link.js";

// const isSessionFoundBoolean = async () => {
//     let resp=false;
//     let sessionId = getCookie('sessionID');
//     if (sessionId) {
//         const response = await fetch(`${linkApi}verifySession`, {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json',
//             },
//             body: JSON.stringify({session: sessionId}),
//         });

//         if (response.status===200 && response.headers.get('content-type').includes('application/json')) {
//             resp=true;
//             console.log(response.json());
//         }
//     }
//     return resp;
// }
const isSessionFoundBoolean = async () => {
    let data;
    let sessionId = getCookie('sessionID');
    if (sessionId) {
        const response = await fetch(`${linkApi}verifySession`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({session: sessionId}),
        });

        if (response.status === 200 && response.headers.get('content-type').includes('application/json')) {                       
             data = await response.json();
        }
    }
    return data ;
}

function getCookie(name) {
    let cookieArr = document.cookie.split(";");
    for(let i = 0; i < cookieArr.length; i++) {
        let cookiePair = cookieArr[i].split("=");
        if(name == cookiePair[0].trim()) {
            return decodeURIComponent(cookiePair[1]);
        }
    }
    return null;
}

export {isSessionFoundBoolean}