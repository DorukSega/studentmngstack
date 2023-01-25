let cofinInput = 0;
window.onload = () => {

    const url = "http://127.0.0.1:8090"

    //TOPBAR
    const topbar = Array.from(document.getElementsByClassName("t_item"))
    topbar.forEach(el =>
        el.onclick = () => {
            if (!el.classList.contains("selected")) {
                document.querySelector(".t_item.selected").classList.remove("selected")
                el.classList.add("selected")
                document.querySelector(".tab.selected").classList.remove("selected")
                document.getElementById(el.getAttribute("ttab")).classList.add("selected")
            }
        }
    )

    //TABLE
    const data_select = document.getElementById("data-select")
    const d_options = [
        "empty",
        "students",
        "teachers",
        "departments",
        "courses",
        "classes",
        "passingstudents"
    ]
    let changes = []
    let keys = []
    const d_head = document.getElementById("d_head")
    const d_body = document.getElementById("d_body")
    data_select.onchange = () => {
        const selected = data_select.selectedIndex

        d_head.innerHTML =
            d_body.innerHTML = "" // clear
        document.getElementById("update").setAttribute("disabled", "true")
        if (selected === 0) {
            d_head.innerHTML =
                d_body.innerHTML = ""
        } else {
            fetch(`${url}/${d_options[selected]}`)
                .then(res => res.json())
                .then(json => {
                    keys = Object.keys(json[0])
                    Object.keys(json[0]).forEach(key => {
                        const col = document.createElement("th")
                        col.innerText = key
                        d_head.append(col)
                    })
                    json.forEach(item => {
                        const row = document.createElement("tr")
                        Object.values(item).forEach((con, i) => {
                            const td = document.createElement("td")
                            if (typeof con === 'object' && con !== null) {
                                if (con.Surname)
                                    td.textContent = con.Name + " " + con.Surname;
                                else if (con.Name)
                                    td.textContent = con.Name;
                                else
                                    td.textContent = con.Number;
                            } else
                                if (typeof con === "boolean")
                                    td.textContent = con === true ? "Yes" : "No";
                                else if (Date.parse(con) && String(con).includes("-"))
                                    td.textContent = new Date(con).toDateString();
                                else {
                                    if (selected < 6 && i != 0)
                                        td.setAttribute("contenteditable", "");
                                    td.textContent = con
                                }


                            row.appendChild(td)
                        })
                        d_body.append(row)
                    });
                })
                .then(() => {
                    //UPDATE
                    Array.from(document.querySelectorAll("td[contenteditable]")).forEach(it => {
                        it.oninput = () => {
                            //
                            let items = Array.from(it.parentElement.childNodes).map(x => x.textContent)

                            changes.push([
                                items[0],
                                keys[items.indexOf(it.textContent)],
                                it.textContent
                            ])
                            console.log(changes[changes.length - 1])

                            document.getElementById("update").removeAttribute("disabled")
                        }
                    })
                })

        }

    }
    data_select.onchange()
    //UPDATE
    document.getElementById("update").onclick = () => {
        const selected = d_options[data_select.selectedIndex]
        let isDone = false
        changes.forEach(ch => {
            fetch(`${url}/update/${selected}?id=${ch[0]}&tag=${ch[1]}&change=${ch[2]}`)
                .then((res) => res.text)
                .then(isDone = true)
        })
        if (isDone) {
            document.getElementById("update").setAttribute("disabled", "")
            alert("Action Completed Succesfully")
            data_select.onchange()
        } else {
            alert("Action Failed")
        }
        changes = []
    }

    //SEARCH
    const search = document.getElementById("search")
    search.value = ""
    search.oninput = () => {
        const input = search.value;
        Array.from(document.querySelectorAll("#d_body tr")).forEach(tr => {
            if (!tr.innerHTML.toLowerCase().match(input.toLowerCase())) {
                if (!tr.classList.contains("hidden"))
                    tr.classList.add("hidden");
            } else {
                tr.classList.remove("hidden")
            }
        })
    }

    //ADD
    const a_options = [
        "student",
        "teacher",
        "department"
    ]
    const add_select = document.getElementById("add-select")
    add_select.onchange = () => {
        const selected = a_options[add_select.selectedIndex]
        document.querySelector(".add-sect.selected").classList.remove("selected")
        document.querySelector(`.add-sect[aaof*="${selected}"]`).classList.add("selected")
    }
    add_select.onchange()
    const submitbtns = Array.from(document.getElementsByClassName("addbtn"))
    submitbtns.forEach(btn => {
        btn.onclick = () => {
            const selected = a_options[add_select.selectedIndex]
            const inputs = Array.from(document.querySelectorAll(".add-sect.selected input[type^='text']"))
            let inarr = inputs.map(el => String(el.value))

            if (inarr.includes("")) {
                alert("Error, empty inputs!")
            } else {
                btn.disabled = true;
                switch (add_select.selectedIndex) {
                    case 0:
                        fetch(`${url}/${selected}/add?id=${inarr[0]}&name=${inarr[1]}&surname=${inarr[2]}&grade=${inarr[3]}&year=${inarr[4]}&depid=${inarr[5]}`)
                            .then(() => alert("Action Completed Succesfully"), btn.disabled = false, data_select.onchange())
                        break;
                    case 1:
                        fetch(`${url}/${selected}/add?id=${inarr[0]}&name=${inarr[1]}&surname=${inarr[2]}&isdean=${inarr[3]}&wage=${inarr[4]}&depid=${inarr[5]}`)
                            .then(() => alert("Action Completed Succesfully"), btn.disabled = false, data_select.onchange())
                        break;
                    case 2:
                        fetch(`${url}/${selected}/add?id=${inarr[0]}&bname=${inarr[1]}&dname=${inarr[2]}`)
                            .then(() => alert("Action Completed Succesfully"), btn.disabled = false, data_select.onchange())
                        break;
                }
            }
        }
    })

    //REMOVE
    const r_options = [
        "student",
        "teacher",
        "department"
    ]
    const rembtn = document.getElementById("rembtn")
    rembtn.onclick = () => {
        const selected = r_options[document.getElementById("rem-select").selectedIndex]
        const reminp = String(document.getElementById("reminp").value)
        if (reminp === "") {
            alert("Error, empty input!")
        } else {
            rembtn.disabled = true;
            fetch(`${url}/remove/${selected}/${reminp}`)
                .then(() => alert("Action Completed Succesfully"), rembtn.disabled = false, data_select.onchange())
        }
    }
    //COURSES OF STUDENT

    const cofin = document.getElementById("cofin")
    cofin.oninput = () => {
        //console.log(cofin.value)
        if (cofinInput != cofin.value) {
            cofinInput = cofin.value
            const c_head = document.getElementById("c_head")
            const c_body = document.getElementById("c_body")
            c_head.innerHTML = c_body.innerHTML = "" // clear

            fetch(`${url}/courses/${cofinInput}`)
                .then(res => res.json())
                .then(json => {
                    Object.keys(json[0]).forEach(key => {
                        const col = document.createElement("th")
                        col.innerText = key
                        c_head.append(col)
                    })
                    json.forEach(item => {
                        const row = document.createElement("tr")
                        Object.values(item).forEach(con => {
                            const td = document.createElement("td")
                            if (typeof con === 'object' && con !== null) {
                                td.textContent = con.Surname ? con.Name + " " + con.Surname : con.Number;
                            } else
                                if (Date.parse(con) && String(con).includes("-"))
                                    td.textContent = new Date(con).toDateString();
                                else
                                    td.textContent = con;

                            row.appendChild(td)
                        })
                        c_body.append(row)
                    });
                })

        }
    }
    cofin.oninput()




}