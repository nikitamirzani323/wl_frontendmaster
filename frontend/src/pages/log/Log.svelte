<script>
    import Home from "../log/Home.svelte";
    import Modal_alert from '../../components/Modal_alert.svelte' 
    export let path_api = ""
    export let font_size = ""
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let isModalNotif = false;
    let msg_error = ""
    let akses_page = false;
    async function initapp() {
        msg_error = ""
        const res = await fetch(path_api+"api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ADMIN-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            msg_error = json.message;
            akses_page = false;
        } else {
            akses_page = true;
            initHome();
        }
        if(msg_error != ""){
            isModalNotif = true;
        }
    }
    async function initHome() {
        const res = await fetch(path_api+"api/log", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            let no = 0;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1
                    let note = record[i]["log_note"];
                    let note_result = note.replace("\n", "<br>");
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["log_id"],
                            home_datetime: record[i]["log_datetime"],
                            home_username: record[i]["log_user"],
                            home_page: record[i]["log_page"],
                            home_tipe: record[i]["log_tipe"],
                            home_note: note_result,
                        },
                    ];
                }
            }
        } else {
            // logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 1000);
    };
    const handleLogout = (e) => {
        logout()
    }
    initapp();
</script>
{#if akses_page == true}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleLogout={handleLogout}
        {font_size}
        {listHome}
        {totalrecord}/>
{/if}
<input type="checkbox" id="my-modal-notiffirst" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notiffirst" 
    modal_title="Information" modal_message="{msg_error}"  />