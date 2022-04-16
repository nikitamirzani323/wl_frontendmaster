<script>
    import Modal_alert from '../../components/Modal_alert.svelte'
    import Home from "../adminrule/Home.svelte";
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
        const res = await fetch(path_api+"api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                page: "ADMINRULE-VIEW",
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
        listHome = [];
        const res = await fetch(path_api+"api/adminrule", {
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
        if (!res.ok) {
			isModalNotif = true;
            msg_error = "Maaf Saat Ini Anda Tidak Bisa Mengakses Halaman Ini"
		}else{
            if (json.status == 200) {
                record = json.record;
                if (record != null) {
                    totalrecord = record.length;
                    for (var i = 0; i < record.length; i++) {
                        listHome = [
                            ...listHome,
                            {
                                home_no: i+1,
                                home_id: record[i]["adminrule_idrule"],
                                home_nmrule: record[i]["adminrule_nmrule"],
                                home_rule: record[i]["adminrule_rule"],
                                home_create: record[i]["adminrule_create"],
                                home_update: record[i]["adminrule_update"],
                            },
                        ];
                    }
                }
            } else {
                isModalNotif = true;
                msg_error = "Maaf Sistem Sedang Mengalami Masalah"
            }
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
        {path_api}
        {font_size}
        {token}
        {listHome}
        {totalrecord}/>
{/if}
<input type="checkbox" id="my-modal-notiffirst" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notiffirst" 
    modal_title="Information" modal_message="{msg_error}"  />