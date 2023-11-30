import {  body, ContentForms } from "../helper/bigBlocContent.js";
body.appendChild(ContentForms)
const contentForms=document.querySelector('.content');

const displayFom=(states)=>{

    contentForms.innerHTML=`<div class="registre-content">
                            <div class="form">
                                <!-- login Form -->
                                ${form1(states)}
                                <!-- Regitre Form -->
                                ${form2(states)}
                            </div>
                                <img src="assets/img/back.png" alt="">
                        </div>
                    `
                    
                   
}

function form1(state) {
    if (state===1) {        
        return `<form class="form-1" method="post">
                        <h1>sign in</h1>
                        <div class="form-group">
                        <input type="text" name="Email"  required/><label>Email or Username</label>
                        </div>
                        <div class="form-group">
                        <input type="password" name="Motdepasse" required/><label>Password</label>
                        </div>
                        <div class="bloc-btn">
                            <input type="submit" value="Sign In" class="submit"> 
                            <div class="spinner" style="display: none;"></div>
                        </div>
                        <div class="row">
                        <p >Not Yet Registered? <span>Sign Up</span></p>
                        </div>
                 </form>
        `
    }
    return ``
}

function form2(state) {
    if (state===2) {   
    return `<form class="form-2" method="post">
                <h1>sign up</h1>
                <div class="form-group">
                    <input type="text" name="FirstName"  required/><label>First Name</label>
                </div>
                <div class="form-group">
                    <input type="text" name="LastName"  required/><label>Last Name</label>
                </div>
                <div class="form-group">
                    <input type="text" name="Username"  required/><label>Nickname</label>
                </div>
                <div class="form-group">
                <input type="email" name="Email"  required/><label>Email </label>
                </div>
                <div class="form-group">
                <input type="date" name="Age"  required/> 
                </div>
                <select name="Gender" id="">
                    <option value="">Select Your Gender</option>
                    <option value="male">Male</option>
                    <option value="female">Female</option>

                </select>

                <div class="form-group">
                <input type="password" name="Motdepasse" required/><label>Password</label>
                </div>
                <div class="form-group">
                    <input type="password" name="Confpassword" required/><label>Confirme Password</label>
                </div>
                <div class="bloc-btn">
                    <input type="submit" value="Sign Up" class="submit"> 
                    <div class="spinner" style="display: none;"></div>
                </div>
                <div class="row">
                <p>Already Registered? <span>Sign In</span></p>
                </div>
            </form>
    `
    }
    return ``
    
}

export {displayFom}