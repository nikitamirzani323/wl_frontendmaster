<script>
    import Home from "../pasaran/Home.svelte";
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
                page: "PASARAN_HOME",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            msg_error = json.message;
            akses_page = false;
        } else {
            initHome();
            akses_page = true;
        }
        if(msg_error != ""){
            isModalNotif = true;
        }
    }
    async function initHome() {
        msg_error = "";
        const res = await fetch(path_api+"api/allpasaran", {
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
                if (record != null) {
                    let status_class = "";
                    let active_class = "";
                    let no = 0;
                    for (var i = 0; i < record.length; i++) {
                        no = ++no;
                        if(record[i]["statuspasaran"] == "ONLINE"){
                            status_class = "bg-[#FFEB3B] text-black"
                        }else{
                            status_class = "bg-[#E91E63] text-white"
                        }
                        if(record[i]["statuspasaranactive"] == "ACTIVE"){
                            active_class = "bg-[#8BC34A] text-black"
                        }else{
                            active_class = "bg-[#E91E63] text-white"
                        }
                        listHome = [
                            ...listHome,
                            {
                                home_no: no,
                                home_id: record[i]["pasaran_idpasarantogel"],
                                home_nama: record[i]["pasaran_nmpasarantogel"],
                                home_tipe: record[i]["pasaran_tipepasaran"],
                                home_url: record[i]["pasaran_urlpasaran"],
                                home_diundi: record[i]["pasaran_pasarandiundi"],
                                home_jamtutup: record[i]["pasaran_jamtutup"],
                                home_jamjadwal: record[i]["pasaran_jamjadwal"],
                                home_jamopen: record[i]["pasaran_jamopen"],
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
 