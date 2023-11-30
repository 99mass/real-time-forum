import { linkApi } from "../helper/api_link.js";
// import { isSessionFound } from "../verify_sessions.js";
import { displayFom } from "../pages/signUpSignIn.js";


// afficher les formulaires
displayFom(1);

const formSignIn=document.querySelector('.form-1');
const formSignUp=document.querySelector('.form-2');
const _ContentForms=document.querySelector('.content');
let spinner = document.querySelector('.spinner');
let rowFormSignIn=document.querySelector('.form-1 .row');
let rowFormSignUp=document.querySelector('.form-2 .row');

if (rowFormSignIn) {    
    rowFormSignIn.addEventListener('click',()=>{
        formSignIn.remove()
        displayFom(2);
    })
}

if (rowFormSignUp) {
    
    rowFormSignUp.addEventListener('click',()=>{
        alert('yes')
        formSignUp.remove();
        displayFom(1);
    })
}


const signUpForm=()=> {

    formSignUp.addEventListener('submit', function(event) {
        event.preventDefault();

         // Show spinner
         spinner.style.display = 'block';

        let firstName = document.querySelector('input[name="FirstName"]').value;
        let lastName = document.querySelector('input[name="LastName"]').value;
        let username = document.querySelector('input[name="Username"]').value;
        let email = document.querySelector('input[name="Email"]').value;
        let age = document.querySelector('input[name="Age"]').value;
        let gender = document.querySelector('select[name="Gender"]').value;
        let password = document.querySelector('input[name="Motdepasse"]').value;
        let confPassword = document.querySelector('input[name="Confpassword"]').value;

        let data = {
            FirstName: firstName,
            LastName: lastName,
            Username: username,
            Email: email,
            Age: age,
            Gender: gender,
            Motdepasse: password,
            ConfPassword: confPassword
        };
        console.log(data);

        fetch(`${linkApi}register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);

            // Hide spinner
            spinner.style.display = 'none';
        })
        .catch((error) => {
            console.error('Error:', error);

            // Hide spinner
            spinner.style.display = 'none';
        });
    });
}

const  signInForm = ()=> {
    
    formSignIn.addEventListener('submit', function(event) {
        event.preventDefault();

        // Show spinner
        spinner.style.display = 'block';

        let email = document.querySelector('input[name="Email"]').value;
        let password = document.querySelector('input[name="Motdepasse"]').value;

        let data = {
            Email: email,
            Motdepasse: password
        };
        console.log(data);
        
        fetch(`${linkApi}signin`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
        .then(response => {
            // Hide spinner
            spinner.style.display = 'none';
           
            if(response.status===200){
               
             if (formSignIn) {
                formSignIn.remove();
            }
            if (formSignUp) {
                formSignUp.remove();
            }
            _ContentForms.remove();
                
              window.location.reload();
            }
            return  response.json();
            
        })
        .then(data => {
            console.log(data);
        })
        .catch((error) => {
             // Hide spinner
             spinner.style.display = 'none';
            console.error('Error:', error);
        });
    });
}

export {signUpForm,signInForm}