

const displayFormMecanisme=(formSignUp,formSignIn,row)=>{
    formSignUp.style.display="none";
    if (row[0]) {    
        row[0].addEventListener('click',()=>{
           
            if (formSignIn) {
                formSignIn.style.display="none";                            
                formSignUp.style.display="block";
            }
        })
    }
    if (row[1]) {   
        row[1].addEventListener('click',()=>{
            if (formSignUp) {
                formSignUp.style.display="none";
                formSignIn.style.display="block";  
            }
        })
    }
}

export{
    displayFormMecanisme
}