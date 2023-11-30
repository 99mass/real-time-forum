import { linkApi } from "./helper/api_link.js";
import { signUpForm, signInForm } from "./auth/forms.js";

import { indexPage } from "./pages/index.js";

const _formSignIn=document.querySelector('.form-1');
const _formSignUp=document.querySelector('.form-2');
const _header=document.querySelector('header');
const _corps=document.querySelector('.corps');
const _ContentForms=document.querySelector('.content');


const isSessionFound = () => {
   
    let sessionId = getCookie('sessionID');
    if (sessionId) {
        console.log("Session found: " + sessionId);
        fetch(`${linkApi}verifySession`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({session: sessionId}),
        })
        .then(response => {
            
            if (response.headers.get('content-type').includes('application/json')) {
                if (_formSignIn) {
                    _formSignIn.remove();
                }
                if (_formSignUp) {
                    _formSignUp.remove();
                }
                _ContentForms.remove()
               

                indexPage();

                return response.json();
            } else {
                if (_header) _header.remove();
                if (_corps) _corps.remove();

                
                signInForm();
                console.log(response.status);
                throw new Error('Received non-JSON response');
            }
        })
        .then(data => console.log(data))
        .catch((error) => {
            if (_header) _header.remove();
            if (_corps) _corps.remove();
            signInForm();

          console.error('Error:', error);
        });
    } else {
        if (_header) _header.remove();
        if (_corps) _corps.remove();
        signInForm();
        console.log("No session found");
    }
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

export {isSessionFound}