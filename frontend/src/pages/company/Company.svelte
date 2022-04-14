<script>
    import Home from "../company/Home.svelte";
    import Modal_alert from '../../components/Modal_alert.svelte' 

    export let path_api = ""
    let listHome = [];
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
                page: "COMPANY_HOME",
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
                company_search: "",
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
                if (record != null) {
                    let home_status_class = "";
                    let home_winlose_class = "text-blue-700";
                    let home_winlosetemp_class = "text-blue-700";
                    let home_selisihwinlose_class = "text-blue-700";
                    for (var i = 0; i < record.length; i++) {
                        let selisihwinlose = parseInt(record[i]["company_winlosetemp"]) - parseInt(record[i]["company_winlose"])
                        if(record[i]["company_status"] == "ACTIVE"){
                            home_status_class = "bg-[#8BC34A] text-black"
                        }else{
                            home_status_class = "bg-red-600 text-white"
                        }
                        if (parseInt(record[i]["company_winlose"]) > 0) {
                            home_winlose_class = "text-blue-700";
                        } else {
                            home_winlose_class = "text-red-500";
                        }
                        if (parseInt(record[i]["company_winlosetemp"]) > 0) {
                            home_winlosetemp_class = "text-blue-700";
                        } else {
                            home_winlosetemp_class = "text-red-500";
                        }
                        if (selisihwinlose > 0) {
                            home_selisihwinlose_class = "text-blue-700";
                        } else {
                            home_selisihwinlose_class = "text-red-500";
                        }
                        listHome = [
                            ...listHome,
                            {
                                home_no: record[i]["company_no"],
                                home_idcompany: record[i]["company_idcompany"],
                                home_startjoin: record[i]["company_startjoin"],
                                home_endjoin: record[i]["company_endjoin"],
                                home_curr: record[i]["company_curr"],
                                home_name: record[i]["company_name"],
                                home_owner: record[i]["company_owner"],
                                home_phone: record[i]["company_phone"],
                                home_email: record[i]["company_email"],
                                home_periode: record[i]["company_periode"],
                                home_winlose: record[i]["company_winlose"],
                                home_winlosetemp: record[i]["company_winlosetemp"],
                                home_winlose_class: home_winlose_class,
                                home_winlosetemp_class: home_winlosetemp_class,
                                home_selisihwinlose: selisihwinlose,
                                home_selisihwinlose_class: home_selisihwinlose_class,
                                home_status: record[i]["company_status"],
                                home_status_class: home_status_class,
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
        {totalrecord}/>
{/if}
<input type="checkbox" id="my-modal-notiffirst" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notiffirst" 
    modal_title="Information" modal_message="{msg_error}"  />