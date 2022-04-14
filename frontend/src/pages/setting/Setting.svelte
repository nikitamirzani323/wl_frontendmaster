<script>
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Input_custom from '../../components/Input.svelte' 
    import Loader from '../../components/Loader.svelte' 
    
    import dayjs from "dayjs";
    export let path_api = ""
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let maintenance_start_field = ""
    let maintenance_end_field = ""
    let akses_page = false;
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let isModalNotif = false;
    let msg_error = ""
    async function initapp() {
        msg_error = ""
        const res = await fetch(path_api+"api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "SETTING_HOME",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            msg_error = json.message;
            akses_page = false;
        } else {
            initSetting();
            akses_page = true;
        }
        if(msg_error != ""){
            isModalNotif = true;
        }
    }
    async function initSetting() {
        const res = await fetch(path_api+"api/setting", {
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
        let record = json.record;
        if (json.status === 400) {
            logout();
        } else {
            
            for (let i = 0; i < record.length; i++) {
                let start = dayjs().format("DD MMM YYYY ")+record[i]["maintenance_start"];
                let end = dayjs().format("DD MMM YYYY ")+record[i]["maintenance_end"];
                maintenance_start_field = dayjs(start).format("HH:mm");
                maintenance_end_field = dayjs(end).format("HH:mm");
            }
        }
    }
    async function saveEntry() {
        let flag = false;
        msg_error = "";
        if (maintenance_start_field == "") {
            flag = true;
            msg_error += "The Jam Maintenance is required<br>";
        }
        if (maintenance_end_field == "") {
            flag = true;
            msg_error += "The Jam Maintenance is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savesetting", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    master: master,
                    maintenance_start: maintenance_start_field,
                    maintenance_end: maintenance_end_field,
                }),
            });
            const json = await res.json();
            if(!res.ok){
                loader_msg = "System Mengalami Trouble"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
            }else{
                if (json.status == 200) {
                    loader_msg = json.message
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                initSetting();
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp();
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />
{#if akses_page}
    <div class=" mx-auto px-2  ">
        <div class="flex gap-2">
            <div class="bg-white shadow-lg w-1/5 ">
                <div class="flex flex-col w-full h-96 ">
                    <div class="border-slate-200">
                        <div class="flex justify-start items-center  border-b-2 p-2">
                            <h1 class="text-slate-600 font-bold text-sm lg:text-2xl uppercase w-full ">Maintenance Game</h1>
                            <div class="flex justify-end w-1/5 p-2 ">
                                <button on:click={() => {
                                    saveEntry();
                                }} class="btn btn-primary hover:bg-primary m-0 p-2 h-1 min-h-[40px] px-3  rounded-md shadow-lg shadow-[#b3e4fc] border-none">Save</button>
                            </div>
                        </div>
                    </div>
                    <div class="w-full mt-2 p-2 flex flex-col gap-5 ">
                        <Input_custom
                            input_required={true}
                            input_tipe="time"
                            bind:value={maintenance_start_field}
                            input_id="maintenance_start_field"
                            input_placeholder="Start Time"/>
                        
                        <Input_custom
                            input_required={true}
                            input_tipe="time"
                            bind:value={maintenance_end_field}
                            input_id="maintenance_end_field"
                            input_placeholder="End Time"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
{/if}

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />