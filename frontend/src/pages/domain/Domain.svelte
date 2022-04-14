<script>
    import Home from "../domain/Home.svelte";
    export let path_api = ""
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master")
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
                page: "DOMAIN_HOME",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch(path_api+"api/domain", {
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
            let no = 0
            let status_class = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if (record[i]["domain_status"] == "RUNNING") {
                        status_class = "bg-[#8BC34A] "
                    } else {
                        status_class = "bg-[#E91E63] text-white"
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["domain_iddomain"],
                            home_nama: record[i]["domain_name"],
                            home_tipe: record[i]["domain_tipe"],
                            home_status: record[i]["domain_status"],
                            home_statusclass: status_class,
                            home_create: record[i]["domain_create"],
                            home_update: record[i]["domain_update"],
                        },
                    ];
                }
            }
        } else {
            logout();
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
        {master}
        {token}
        {listHome}
        {totalrecord}/>
{/if}