import { routes, replaceRouter, currentPath} from "../router/route.js";


const formAddPost=(formCreatPost)=>{
    formCreatPost.addEventListener('submit', function(e) {
        e.preventDefault();
        let errorPostForm=document.querySelector('.error-post-form');
       
        let selectedCategories = [];
        document.querySelectorAll('input[name="categories"]:checked').forEach((checkbox) => {
            selectedCategories.push(checkbox.value);
        });

        let title = document.querySelector('input[name="Title"]').value.trim();
        let content = document.querySelector('textarea[name="Content"]').value.trim();
        let file = document.querySelector('input[type="file"]').files[0];

        if (file) {
            let fileType = file.type;
            let fileSize = file.size / (1024 * 1024); // size in MB

            if (!(['image/jpeg', 'image/svg+xml', 'image/png', 'image/gif'].includes(fileType) && fileSize <= 20)) {
                errorPostForm.innerHTML='File format is not valid';
                return;
            }
        }
        let reader = new FileReader();
        reader.onloadend = function() {
            let base64File = reader.result;
            let data = {
                Title: title,
                Content: content,
                Category: selectedCategories 
            };
        
            if (file) {
                data.image = base64File;
            }
        
            fetch('http://localhost:8080/addpost', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(response =>{ 
                if (response.status===200) {
                    let r = routes["/Home"]['name'];
                    replaceRouter(r);
                    window.location.reload();
                }
                return  response.json()
            })
            .then(data =>{ 
                if (data["message"]) {                    
                    // alert(data["message"]);
                    errorPostForm.innerHTML=data["message"];
                }
            })
            .catch((error) => {
                console.error('Error:', error);
            });
        }
        
        if (file) {
            reader.readAsDataURL(file);
        } else {
            reader.onloadend();
        }
    });


}

export{formAddPost}
