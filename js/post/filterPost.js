import { linkApi } from "../helper/api_link.js";


const filterPost=(contenCatId,categoryID,callback)=>{

   for (let i = 0; i < contenCatId.length; i++) {
        const cat = contenCatId[i];
        cat.addEventListener("click", function () {

        let _categoryID = categoryID[i].textContent.trim();
        fetch(`${linkApi}category`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ CategoryID: _categoryID })
            })
            .then(response => {
                if (response.status !== 200) {
                    console.log('Error:' + response["message"]);
                    return;
                }
                return response.json();
            })
            .then(data => { 
                callback(data);
            })
            .catch((error) => {
                console.error('Error   : no comment found');
            });
        });
    }

}



export{
    filterPost
}