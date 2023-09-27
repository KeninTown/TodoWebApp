
//delete todo
deleteButtons = document.getElementsByClassName("delete-button")

function deleteTodo(id, nodeToDelete){
    return async function() {
        await fetch("http://localhost/todos", {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({id: id})
        })

        nodeToDelete.parentNode.removeChild(nodeToDelete)
        return false
    }
}

for(let i = 0; i < deleteButtons.length; i++){
    let todo = deleteButtons[i] 
    let todoId = todo.getAttribute("data-todo-id")
    let removeTodo = document.getElementById(`card-${todoId}`)


    todo.onclick = deleteTodo(Number(todoId), removeTodo)
}




//done todo
doneButtons = document.getElementsByClassName("done-button")

function doneTodo(id, btn, todo){
    return async function() {
        if (btn.checked){
            await fetch("http://localhost/complete", {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({id: id, completed: true})
            })

            todo.classList.add("done")

        } else{
            await fetch("http://localhost/complete", {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({id: id, completed: false})
            })
            todo.classList.remove("done")
        }
        return false
    }
}

for(let i = 0; i < doneButtons.length; i++){
    let doneTodoBtn = doneButtons[i] 
    let todoId = doneTodoBtn.getAttribute("data-todo-id")
    let todoToDone = document.getElementById(`card-${todoId}`)
    
    doneTodoBtn.onchange = doneTodo(Number(todoId), doneTodoBtn, todoToDone)
}

hiddenBtn = document.getElementById("hideTodo")

hiddenBtn.onchange = function(){
    if (hiddenBtn.checked){
        doneTodos = document.querySelectorAll(".card.done")
        for(let i = 0; i < doneTodos.length; i++){
            doneTodos[i].classList.add("hidden")
        }
    } else{
        hiddenTodos = document.querySelectorAll(".card.hidden")
        for(let i = 0; i < hiddenTodos.length; i++){
            hiddenTodos[i].classList.remove("hidden")
        }
    }
    return false
}


function chooseTag(tag){
    return function() {
        let labelId = tag.getAttribute("id")
        let label = document.querySelector(`label[for="${labelId}"]`)
        if (tag.checked){
            label.classList.add("choosed")

        } else{
            label.classList.remove("choosed")
        }
        return false
    }
}

let tags = document.getElementsByClassName("tag_btn")

for(let i = 0; i < tags.length; i++){
    let chosedTag = tags[i] 
    chosedTag.onchange = chooseTag(chosedTag)
}

let cancel = document.getElementById("cancel_modal")

cancel.onclick = function(){
    let modal = document.getElementById("modal")
    modal.classList.add("hidden")
    cleanModal()
}

let create = document.getElementById("open_modal")

create.onclick = function(){
    let modal = document.getElementById("modal")
    modal.classList.remove("hidden")

    let submit = document.getElementById("submit")

    submit.onclick = async function(){
        let title = document.getElementById("header").value

        if (!title){
            alert("Введите заголовок")
            return
        }

        let description = document.getElementById("description").value

        if (!description){
            alert("Введите описание")
            return
        }

        let tags = document.getElementsByClassName("tag_btn")
        
        let type = []

        for(let i = 0; i < tags.length; i++){
            if (tags[i].checked){
                type.push(tags[i].getAttribute("name"))
            }
        }
        if(type.length === 0){
            alert("Выберите тег")
            return
        }
        let response = await fetch("http://localhost/todos", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        title,
                        type,
                        description
                    })
        })

        let todo = await response.json()

        tagsImg = ``
        for(let i = 0; i < todo.type.length; i++){
            tagsImg = `${tagsImg}<img src="/static/imgs/${todo.type[i]}.svg" alt="${todo.type[i]}">`
        }
        todoHtml = `
        <div class="card" id="card-${todo.id}">
            <div class="header_card">
                <div class="title">
                    ${todo.title}
                </div>
                <div class="dots"> 
                    <input type="image" src="/static/imgs/dots.svg" alt="Кнопка dots" id="dots"> 
                    <ul class="dropdown_content">
                        <li><a href="#" data-todo-id="${todo.id}" class="delete-button">Удалить</a></li>
                        <li><a href="#" data-todo-id="${todo.id}">Редактировать</a></li>
                    </ul>
                </div>
            </div>
            <div class="article">${todo.description}</div>
            <div class="footer_card">
                <div class="card-tags">
                    ${tagsImg}
                </div>
                <div class="done_checkbox_container">
                    <label class="done">
                        <input type="checkbox" name="done_checkbox" class="done-button" data-todo-id="${todo.id}">
                        <span class="done-checkbox"></span>
                    Выполнено
                    </label>
                </div>
            </div>
        </div>`
        let leftColumn = document.querySelector("div.left_column")
        let rightColumn = document.querySelector("div.right_column")

        if(leftColumn.children.length >= rightColumn.children.length){
            leftColumn.insertAdjacentHTML("beforeend", todoHtml)
        }


        //done todo
        let doneTodoBtn = document.querySelector(`input.done-button[data-todo-id="${todo.id}"]`)
        let todoToDone = document.getElementById(`card-${todo.id}`)
        doneTodoBtn.onchange = doneTodo(Number(todo.id), doneTodoBtn, todoToDone)

        //delete todo
        let deletTodoBtn = document.querySelector(`a.delete-button[data-todo-id="${todo.id}"]`)
        let removeTodo = document.getElementById(`card-${todo.id}`)
        deletTodoBtn.onclick = deleteTodo(Number(todo.id), removeTodo)

        let updateTodoBtn = document.querySelector(`a.delete-button[data-todo-id="${todo.id}"]`)
        let updateTodo = document.getElementById(`card-${todo.id}`)
        updateTodoBtn.onclick = openUpdateModal(todo)

        cancel.click()
    }
}

function cleanModal(){
    header = document.getElementById("header")
    header.value = ""

    todoBody = document.getElementById("description")
    todoBody.value = ""

    for(let i = 0; i < tags.length; i++){
        tags[i].checked = false
        chooseTag(tags[i])()
    }
}




//update todo
updateBtns = document.getElementsByClassName("update-button")

function openUpdateModal(currentTodo){
    return async function() {

        let modal = document.getElementById("modal") 
        modal.classList.remove("hidden")

        let header = document.getElementById("header")
        header.value = currentTodo.title

        let description = document.getElementById("description")
        description.value = currentTodo.description

        for(let i = 0; i < currentTodo.type.length; i++){
            let type = currentTodo.type[i] 
            let tag = document.getElementById(`${type}_tag`) 
            tag.checked = true
            let label = document.querySelector(`label[for="${type}_tag"]`)
            label.classList.add("choosed")
        }

        let submitBtn = document.getElementById("submit")


        submitBtn.onclick = updateTodoFunc(currentTodo.id)
        return false
    }
}

function updateTodoFunc(cardId){
    return async function(){
                
        let title = document.getElementById("header").value

        if (!title){
            alert("Введите заголовок")
            return
        }
    
        let description = document.getElementById("description").value
    
        if (!description){
            alert("Введите описание")
            return
        }
    
        let tags = document.getElementsByClassName("tag_btn")
        
        let type = []
    
        for(let i = 0; i < tags.length; i++){
            if (tags[i].checked){
                type.push(tags[i].getAttribute("name"))
            }
        }
        if(type.length === 0){
            alert("Выберите тег")
            return
        }
        
        let response = await fetch("http://localhost/todos", {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        id: Number(cardId),
                        title,
                        type,
                        description
                    })
        })
    
        let todo = await response.json()

        document.getElementById(`title-${todo.id}`).textContent = todo.title
        document.getElementById(`article-${todo.id}`).textContent = todo.description


        let updatedTags = document.getElementById(`tags-${todo.id}`)
        let imgs = ``
        types = todo.type
        for(let i = 0; i < types.length; i++){
            imgs = `${imgs}<img src="/static/imgs/${types[i]}.svg" alt="${types[i]}">`
        }
        updatedTags.innerHTML = ""
        updatedTags.innerHTML = imgs
        cancel.click()
    } 
}

for(let i = 0; i < updateBtns.length; i++){
    let btn = updateBtns[i] 
    let todoId = btn.getAttribute("data-todo-id")
    let todoToUpdate = document.getElementById(`card-${todoId}`)

    let title = document.getElementById(`title-${todoId}`).textContent
    let description = document.getElementById(`article-${todoId}`).textContent

    let tagsContainer = document.getElementById(`tags-${todoId}`)
    let tags =  tagsContainer.children
    types = []
    for(let i = 0; i < tags.length; i++){
        tag = tags[i]
        types.push(tag.getAttribute("alt"))
    }

    let currentTodo = {
        id: todoId,
        title,
        description,
        type: types
    } 
    btn.onclick = openUpdateModal(currentTodo, todoToUpdate)
}