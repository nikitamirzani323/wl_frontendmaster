<script>
    import Home from "../company/Home.svelte";
    import Modal_alert from '../../components/Modal_alert.svelte' 

    export let path_api = ""
    let listHome = [];
    let listcurrency = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let akses_page = false;
    let isModalNotif = false;
    let msg_error = ""
    async function initapp() {
        const res = await fetch(path_api+"api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                page: "COMPANY-VIEW",
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
        const res = await fetch(path_api+"api/company", {
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
                totalrecord = record.length;
                let recordlistcurr = json.listcurrency;
                if (record != null) {
                    let home_status_text = "";
                    let home_status_class = "";
                    for (var i = 0; i < record.length; i++) {
                        if(record[i]["company_status"] == "Y"){
                            home_status_class = "bg-[#8BC34A] text-black"
                            home_status_text = "ACTIVE"
                        }else{
                            home_status_class = "bg-red-600 text-white"
                            home_status_text = "DEACTIVE"
                        }
                        listHome = [
                            ...listHome,
                            {
                                home_no: i+1,
                                home_idcompany: record[i]["company_idcompany"],
                                home_startjoin: record[i]["company_startjoin"],
                                home_endjoin: record[i]["company_endjoin"],
                                home_curr: record[i]["company_idcurr"],
                                home_name: record[i]["company_nmcompany"],
                                home_owner: record[i]["company_nmowner"],
                                home_phone: record[i]["company_phoneowner"],
                                home_email: record[i]["company_emailowner"],
                                home_urlendpoint: record[i]["company_urlendpoint"],
                                home_create: record[i]["company_create"],
                                home_update: record[i]["company_update"],
                                home_status: home_status_text,
                                home_status_class: home_status_class,
                                
                            },
                        ];
                    }
                }
                if (recordlistcurr != null) {
                    for (let i = 0; i < recordlistcurr.length; i++) {
                        listcurrency = [
                            ...listcurrency,
                            {
                                curr_idcurr:recordlistcurr[i]["curr_idcurr"],
                            },
                        ];
                    }
                }
            } else {
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
        {token}
        {master}
        {listHome}
        {listcurrency}
        {totalrecord}/>
{/if}
<input type="checkbox" id="my-modal-notiffirst" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notiffirst" 
    modal_title="Information" modal_message="{msg_error}"  />