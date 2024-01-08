import {  body, ContentForms } from "../helper/bigBlocContent.js";
import {  routes,addRouter,replaceRouter,currentPath } from "../router/route.js";


const displayFom=()=>{

    body.appendChild(ContentForms)
    const contentForms=document.querySelector('.content');

    contentForms.innerHTML=`<div class="registre-content">
                            <div class="form">
                                <!-- login Form -->
                                ${form1()}
                                <!-- Regitre Form -->
                                ${form2()}
                            </div>
                                <img src="assets/img/back.png" alt="">
                        </div>
    `;
           
}


function form1() {
    let r=routes["/Login"]['name'];
    replaceRouter(r);
    const form = document.createElement('form');
    form.className = 'form-1';
    form.method = 'post';

    const h1 = document.createElement('h1');
    h1.textContent = 'sign in';
    form.appendChild(h1);

    const err= document.createElement('p');
    err.className="error-page";
    err.textContent = 'mass';
    err.style.textAlign="center";
    err.style.margin="5px 0px";
    err.style.color="white";
    err.style.background="indianred";
    err.style.padding="5px 0px";
    err.style.fontSize="14px";
    err.style.fontWeight="initial";
    err.style.display="none"
    form.appendChild(err);


    const formElements = [
        { type: 'text', name: 'Email', label: 'Email or Nickname' },
        { type: 'password', name: 'Motdepasse', label: 'Password' }
    ];

    formElements.forEach(el => {
        const div = document.createElement('div');
        div.className = 'form-group';

        const input = document.createElement('input');
        input.type = el.type;
        input.name = el.name;
        input.required = true;
        input.setAttribute("autocomplete", "off");

        const label = document.createElement('label');
        label.textContent = el.label;

        div.appendChild(input);
        div.appendChild(label);
        form.appendChild(div);
    });

    const divBtn = document.createElement('div');
    divBtn.className = 'bloc-btn';

    const inputSubmit = document.createElement('input');
    inputSubmit.type = 'submit';
    inputSubmit.value = 'Sign In';
    inputSubmit.className = 'submit';

    const divSpinner = document.createElement('div');
    divSpinner.className = 'spinner';
    divSpinner.style.display = 'none';

    divBtn.appendChild(inputSubmit);
    divBtn.appendChild(divSpinner);
    form.appendChild(divBtn);

    const divRow = document.createElement('div');
    divRow.className = 'row';

    const p = document.createElement('p');
    p.innerHTML = 'Not Yet Registered? <span>Sign Up</span>';

    divRow.appendChild(p);
    form.appendChild(divRow);

    return form.outerHTML;
}

function form2() {
    
    const form = document.createElement('form');
    form.className = 'form-2';
    form.method = 'post';

    const h1 = document.createElement('h1');
    h1.textContent = 'sign up';
    form.appendChild(h1);

    const err= document.createElement('p');
    err.className="error-page";
    err.textContent = 'mass';
    err.style.textAlign="center";
    err.style.margin="5px 0px";
    err.style.color="white";
    err.style.background="indianred";
    err.style.padding="5px 0px";
    err.style.fontSize="14px";
    err.style.fontWeight="initial";
    err.style.display="none"
    form.appendChild(err);

    const formElements = [
        { type: 'text', name: 'FirstName', label: 'First Name' },
        { type: 'text', name: 'LastName', label: 'Last Name' },
        { type: 'text', name: 'Username', label: 'Nickname' },
        { type: 'email', name: 'Email', label: 'Email' },
        { type: 'text', name: 'Age', label: 'Age' },
        { type: 'password', name: 'Motdepasse', label: 'Password' },
        { type: 'password', name: 'Confpassword', label: 'Confirm Password' }
    ];

    formElements.forEach(el => {
        const div = document.createElement('div');
        div.className = 'form-group';

        const input = document.createElement('input');
        input.type = el.type;
        input.name = el.name;
        input.required = true;
        input.setAttribute("autocomplete", "off");

        const label = document.createElement('label');
        label.textContent = el.label;

        div.appendChild(input);
        div.appendChild(label);
        form.appendChild(div);
    });

    const select = document.createElement('select');
    select.name = 'Gender';
    select.id = '';
    ['Select Your Gender', 'Male', 'Female'].forEach((optionText, i) => {
        const option = document.createElement('option');
        option.value = i === 0 ? '' : optionText.toLowerCase();
        option.textContent = optionText;
        select.appendChild(option);
    });
    form.appendChild(select);

    const divBtn = document.createElement('div');
    divBtn.className = 'bloc-btn';

    const inputSubmit = document.createElement('input');
    inputSubmit.type = 'submit';
    inputSubmit.value = 'Sign Up';
    inputSubmit.className = 'submit';

    const divSpinner = document.createElement('div');
    divSpinner.className = 'spinner';
    divSpinner.style.display = 'none';

    divBtn.appendChild(inputSubmit);
    divBtn.appendChild(divSpinner);
    form.appendChild(divBtn);

    const divRow = document.createElement('div');
    divRow.className = 'row';

    const p = document.createElement('p');
    p.innerHTML = 'Already Registered? <span>Sign In</span>';

    divRow.appendChild(p);
    form.appendChild(divRow);

    return form.outerHTML;
}


export {displayFom}