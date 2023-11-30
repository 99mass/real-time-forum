import { linkApi } from "../helper/api_link.js";

const formSignIn=document.querySelector('.form-1');
const formSignUp=document.querySelector('.form-2');
let spinner = document.querySelector('.spinner');
let rowFormSignIn=document.querySelector('.form-1 .row');
let rowFormSignUp=document.querySelector('.form-2 .row');

formSignUp.style.display="none"
rowFormSignIn.addEventListener('click',()=>{
    if (formSignUp.style.display==="none") {
        formSignUp.style.display="block"
        formSignIn.style.display="none"
    }
})
rowFormSignUp.addEventListener('click',()=>{
    if (formSignIn.style.display==="none") {
        formSignIn.style.display="block"
        formSignUp.style.display="none"
    }
})

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
            Password: password,
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
        .then(response => response.json())
        .then(data => {
             // Hide spinner
             spinner.style.display = 'none';
            console.log('Success:', data);
        })
        .catch((error) => {
             // Hide spinner
             spinner.style.display = 'none';
            console.error('Error:', error);
        });
    });
}

export {signUpForm,signInForm}