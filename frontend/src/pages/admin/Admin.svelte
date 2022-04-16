<script>
    import Home from "../admin/Home.svelte";
    import Modal_alert from '../../components/Modal_alert.svelte' 
    export let path_api = ""
    export let font_size = "";
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let akses_page = false;
    let isModalNotif = false;
    let msg_error = ""
    let admin_listrule = [];
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
        listHome = []
        admin_listrule = []
        const res = await fetch(path_api+"api/admin", {
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
                let recordlistrule = json.listadminrule;
                let status_class = "";
                if (record != null) {
                    totalrecord = record.length;
                    for (var i = 0; i < record.length; i++) {
                        if (record[i]["admin_status"] == "ACTIVE") {
                            status_class = "bg-[#ebfbee] text-[#6ec07b]"
                        } else {
                            status_class = "bg-[#fde3e3] text-[#ea7779]"
                        }
                        listHome = [
                            ...listHome,
                            {
                                home_no: i + 1,
                                home_username: record[i]["admin_username"],
                                home_nama: record[i]["admin_nama"],
                                home_email: record[i]["admin_email"],
                                home_phone: record[i]["admin_phone"],
                                home_idrule: record[i]["admin_idrule"],
                                home_rule: record[i]["admin_rule"],
                                home_joindate: record[i]["admin_joindate"],
                                home_lastlogin: record[i]["admin_lastlogin"],
                                home_lastipaddres: record[i]["admin_lastipaddres"],
                                home_status: record[i]["admin_status"],
                                home_create: record[i]["admin_create"],
                                home_update: record[i]["admin_update"],
                                home_statusclass: status_class,
                            },
                        ];
                    }
                }
                if (recordlistrule != null) {
                    for (let i = 0; i < recordlistrule.length; i++) {
                        admin_listrule = [
                            ...admin_listrule,
                            {
                                adminrule_idrule:recordlistrule[i]["adminrule_idrule"],
                                adminrule_nmrule:recordlistrule[i]["adminrule_nmrule"].toUpperCase(),
                            },
                        ];
                    }
                }
            }else {
                isModalNotif = true;
                msg_error = "Maaf Sistem Sedang Mengalami Masalah"
            }
        }
        setTimeout(function () {
            isModalNotif = false;
        }, 1000);
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        admin_listrule = [];
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
        {admin_listrule}
        {listHome}
        {totalrecord}/>
{/if}
<input type="checkbox" id="my-modal-notiffirst" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notiffirst" 
    modal_title="Information" modal_message="{msg_error}"  />