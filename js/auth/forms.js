
import { isSessionFoundBoolean } from "../verify_sessions.js";
import { indexPage } from "../pages/index.js";



const signInForm = async (_ContentForms, formSignIn, formSignUp, spinner, linkApi) => {
    if (formSignIn)
        formSignIn.addEventListener('submit', async function (event) {
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
            const errorp = document.querySelector('.error-page');//pour afficher les erreurs
            try {
                const response = await fetch(`${linkApi}signin`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(data),
                });

                spinner.style.display = 'none'; // Hide spinner

                if (response.status === 200) {
                    async function checkUserSession() {
                        const data = await isSessionFoundBoolean();
                        
                        if (data != undefined && data['Session']) {
                            if (formSignIn) { formSignIn.remove(); }
                            if (formSignUp) { formSignUp.remove(); }
                           

                            _ContentForms.remove();
                            window.location.reload();
                            
                            indexPage(data);
                        } else {
                            errorp.style.display = "block";
                            errorp.innerHTML = "error : veillez tenter encore.";
                            console.log('Error : api inaccessible');
                        }
                    }
                    (async () => {
                        await checkUserSession();
                    })();
                } else {

                    let error = await response.json();
                    errorp.style.display = "block";
                    errorp.innerHTML = error["message"];
                    console.log(error["message"]);

                }


            } catch (error) {
                spinner.style.display = 'none'; // Hide spinner
                errorp.style.display = "block";
                errorp.innerHTML = "error : veillez tenter encore.";
                console.error('Error:', error);
            }
        });
}



const signUpForm = async (_ContentForms, formSignUp, formSignIn, spinner, linkApi) => {
    if (formSignUp)
        formSignUp.addEventListener('submit', async function (e) {
            e.preventDefault();
            spinner.style.display = 'block'; // Show spinner

            let _firstName = document.querySelector('input[name="FirstName"]').value;
            let _lastName = document.querySelector('input[name="LastName"]').value;
            let _username = document.querySelector('input[name="Username"]').value;
            let _gender = document.querySelector('select[name="Gender"]').value;
            let _age = document.querySelector('input[name="Age"]').value;
            let _email = document.querySelectorAll('input[name="Email"]')[1].value;
            let _password = document.querySelectorAll('input[name="Motdepasse"]')[1].value;
            let _confPassword = document.querySelector('input[name="Confpassword"]').value;

            let data = {
                userName: _username,
                firstName: _firstName,
                lastName: _lastName,
                gender: _gender,
                age: _age,
                email: _email,
                password: _password,
                confpassword: _confPassword
            };
            console.log(data);

            const errorp = document.querySelectorAll('.error-page')[1];//pour afficher les erreurs
            try {
                const response = await fetch(`${linkApi}register`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(data),
                })
                spinner.style.display = 'none'; // Hide spinner

                if (response.status === 200) {
                    async function checkUserSession() {
                        const data = await isSessionFoundBoolean();
                        if (data != undefined && data['Session']) {
                            if (formSignUp) { formSignUp.remove(); }
                            if (formSignIn) { formSignIn.remove(); }
                            _ContentForms.remove();
                            window.location.reload();
                            indexPage(data);
                        } else {
                            errorp.style.display = "block";
                            errorp.innerHTML = "error : veillez tenter encore.";
                            console.log('Error : api inaccessible');
                        }
                    }
                    (async () => {
                        await checkUserSession();
                    })();
                } else {
                    let error = await response.json();
                    errorp.style.display = "block";
                    errorp.innerHTML = error["message"];
                    console.log(error["message"]);

                }
            } catch (error) {
                spinner.style.display = 'none'; // Hide spinner
                errorp.style.display = "block";
                errorp.innerHTML = "error : veillez tenter encore.";
                console.error('Error:', error);
            }
        });
}

export { signUpForm, signInForm }