<script>
    import { createEventDispatcher } from "svelte";
    import Input_custom from '../../components/Input.svelte'
    
    export let path_api = "";
    export let master = "";
    export let token = "";
    export let idcompany = "";
    export let companypasaran_id = "";
    export let pasaran_id_field = "";
    export let pasaran_minbet_5050umum_field = 0;
    export let pasaran_maxbet_5050umum_field = 0;
    export let pasaran_keibesar_5050umum_field = 0;
    export let pasaran_keikecil_5050umum_field = 0;
    export let pasaran_keigenap_5050umum_field = 0;
    export let pasaran_keiganjil_5050umum_field = 0;
    export let pasaran_keitengah_5050umum_field = 0;
    export let pasaran_keitepi_5050umum_field = 0;
    export let pasaran_discbesar_5050umum_field = 0;
    export let pasaran_disckecil_5050umum_field = 0;
    export let pasaran_discgenap_5050umum_field = 0;
    export let pasaran_discganjil_5050umum_field = 0;
    export let pasaran_disctengah_5050umum_field = 0;
    export let pasaran_disctepi_5050umum_field = 0;
    export let pasaran_limitglobal_5050umum_field = 0;
    export let pasaran_limittotal_5050umum_field = 0;
    let buttonLoading_class = "btn btn-primary";
    let buttonLoadingfetch_class = "btn btn-warning";
    let msg_error = "";
    let dispatch = createEventDispatcher();
    async function save432d() {
        let flag = false;
        msg_error = "";
        if (pasaran_minbet_5050umum_field == "") {
            flag = true;
            msg_error += "The Min Bet is required<br>";
        }
        if (pasaran_maxbet_5050umum_field == "") {
            flag = true;
            msg_error += "The Max Bet is required<br>";
        }
        if (pasaran_limittotal_5050umum_field == "") {
            flag = true;
            msg_error += "The Limit Total is required<br>";
        }
        if (pasaran_limitglobal_5050umum_field == "") {
            flag = true;
            msg_error += "The Limit Global is required<br>";
        }
        if (pasaran_keibesar_5050umum_field == "") {
            flag = true;
            msg_error += "The Kei Besar is required<br>";
        }
        if (pasaran_keikecil_5050umum_field == "") {
            flag = true;
            msg_error += "The Kei Kecil is required<br>";
        }
        if (pasaran_keigenap_5050umum_field == "") {
            flag = true;
            msg_error += "The Kei Genap is required<br>";
        }
        if (pasaran_keiganjil_5050umum_field == "") {
            flag = true;
            msg_error += "The Kei Ganjil is required<br>";
        }
        if (pasaran_keitengah_5050umum_field == "") {
            flag = true;
            msg_error += "The Kei Tengah is required<br>";
        }
        if (pasaran_keitepi_5050umum_field == "") {
            flag = true;
            msg_error += "The Kei Tepi is required<br>";
        }
        if (pasaran_discbesar_5050umum_field == "") {
            flag = true;
            msg_error += "The Disc Besar is required<br>";
        }
        if (pasaran_disckecil_5050umum_field == "") {
            flag = true;
            msg_error += "The Disc Kecil is required<br>";
        }
        if (pasaran_discgenap_5050umum_field == "") {
            flag = true;
            msg_error += "The Disc Genap is required<br>";
        }
        if (pasaran_discganjil_5050umum_field == "") {
            flag = true;
            msg_error += "The Disc Ganjil is required<br>";
        }
        if (pasaran_disctengah_5050umum_field == "") {
            flag = true;
            msg_error += "The Disc Tengah is required<br>";
        }
        if (pasaran_disctepi_5050umum_field == "") {
            flag = true;
            msg_error += "The Disc Tepi is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/updatecompanypasaran5050umum", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "Edit",
                    master: master,
                    company: idcompany,
                    companypasaran_id: companypasaran_id,
                    pasaran_id: pasaran_id_field,
                    pasaran_minbet_5050umum: parseInt(pasaran_minbet_5050umum_field),
                    pasaran_maxbet_5050umum: parseInt(pasaran_maxbet_5050umum_field),
                    pasaran_limittotal_5050umum: parseInt(pasaran_limittotal_5050umum_field),
                    pasaran_limitglobal_5050umum: parseInt(pasaran_limitglobal_5050umum_field),
                    pasaran_keibesar_5050umum: parseFloat((pasaran_keibesar_5050umum_field / 100).toPrecision(3)),
                    pasaran_keikecil_5050umum: parseFloat((pasaran_keikecil_5050umum_field / 100).toPrecision(3)),
                    pasaran_keigenap_5050umum: parseFloat((pasaran_keigenap_5050umum_field / 100).toPrecision(3)),
                    pasaran_keiganjil_5050umum: parseFloat((pasaran_keiganjil_5050umum_field / 100).toPrecision(3)),
                    pasaran_keitengah_5050umum: parseFloat((pasaran_keitengah_5050umum_field / 100).toPrecision(3)),
                    pasaran_keitepi_5050umum: parseFloat((pasaran_keitepi_5050umum_field / 100).toPrecision(3)),
                    pasaran_discbesar_5050umum: parseFloat((pasaran_discbesar_5050umum_field / 100).toPrecision(3)),
                    pasaran_disckecil_5050umum: parseFloat((pasaran_disckecil_5050umum_field / 100).toPrecision(3)),
                    pasaran_discgenap_5050umum: parseFloat((pasaran_discgenap_5050umum_field / 100).toPrecision(3)),
                    pasaran_discganjil_5050umum: parseFloat((pasaran_discganjil_5050umum_field / 100).toPrecision(3)),
                    pasaran_disctengah_5050umum: parseFloat((pasaran_disctengah_5050umum_field / 100).toPrecision(3)),
                    pasaran_disctepi_5050umum: parseFloat((pasaran_disctepi_5050umum_field / 100).toPrecision(3)),
                }),
            });
            const json = await res.json();
            if(!res.ok){
                let temp_msg = "System Mengalami Trouble"
                dispatch("handleLoadingRunningFinish", {
                        temp_msg
                });
            }else{
                let temp_msg = json.message
                dispatch("handleLoadingRunningFinish", {
                        temp_msg
                });
            }
            buttonLoading_class = "btn btn-primary"
        } else {
            if(msg_error != ""){
                let temp_msg = msg_error
                dispatch("handleCallNotif", {
                        temp_msg
                });
            }
        }
    }
    async function fetchcolok() {
        let flag = false;
        msg_error = "";
        if (pasaran_id_field == "") {
            flag = true;
            msg_error += "The Pasaran is required<br>";
        }
        if (flag == false) {
            buttonLoadingfetch_class = "btn loading"
            dispatch("handleLoadingRunning", "call");
            const res = await fetch(path_api+"api/fetchpasaran5050umum", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "Edit",
                    master: master,
                    company: idcompany,
                    pasaran_id: pasaran_id_field,
                    Companypasaran_id: companypasaran_id,
                }),
            });
            const json = await res.json();
            if(!res.ok){
                let temp_msg = "System Mengalami Trouble"
                dispatch("handleLoadingRunningFinish", {
                        temp_msg
                });
            }else{
                let temp_msg = json.message
                dispatch("handleLoadingRunningFinish", {
                        temp_msg
                });
            }
            buttonLoadingfetch_class = "btn btn-warning"
        } else {
            if(msg_error != ""){
                let temp_msg = msg_error
                dispatch("handleCallNotif", {
                        temp_msg
                });
            }
        }
    }
</script>

<div class="grid grid-cols-4 gap-1 mt-2 mb-5">
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_minbet_5050umum_field}
        input_id="pasaran_minbet_5050umum_field"
        input_placeholder="Minimal Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_limittotal_5050umum_field}
        input_id="pasaran_limittotal_5050umum_field"
        input_placeholder="Limit Total"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keibesar_5050umum_field}
        input_precision=2
        input_id="pasaran_keibesar_5050umum_field"
        input_placeholder="KEI BESAR"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_discbesar_5050umum_field}
        input_precision=2
        input_id="pasaran_discbesar_5050umum_field"
        input_placeholder="DISC BESAR(%)"/>
    
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_maxbet_5050umum_field}
        input_id="pasaran_maxbet_5050umum_field"
        input_placeholder="Max Bet"/>
    <Input_custom
        input_enabled={true}
        input_tipe="number"
        input_maxlenght="12"
        bind:value={pasaran_limitglobal_5050umum_field}
        input_id="pasaran_limitglobal_5050umum_field"
        input_placeholder="Limit Total"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keikecil_5050umum_field}
        input_precision=2
        input_id="pasaran_keikecil_5050umum_field"
        input_placeholder="KEI KECIL"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disckecil_5050umum_field}
        input_precision=2
        input_id="pasaran_disckecil_5050umum_field"
        input_placeholder="DISC KECIL(%)"/>
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keigenap_5050umum_field}
        input_precision=2
        input_id="pasaran_keigenap_5050umum_field"
        input_placeholder="KEI GENAP"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_discgenap_5050umum_field}
        input_precision=2
        input_id="pasaran_discgenap_5050umum_field"
        input_placeholder="DISC GENAP(%)"/>
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keiganjil_5050umum_field}
        input_precision=2
        input_id="pasaran_keiganjil_5050umum_field"
        input_placeholder="KEI GANJIL"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_discganjil_5050umum_field}
        input_precision=2
        input_id="pasaran_discganjil_5050umum_field"
        input_placeholder="DISC GANJIL(%)"/>
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keitengah_5050umum_field}
        input_precision=2
        input_id="pasaran_keitengah_5050umum_field"
        input_placeholder="KEI TENGAH"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disctengah_5050umum_field}
        input_precision=2
        input_id="pasaran_disctengah_5050umum_field"
        input_placeholder="DISC TENGAH(%)"/>
    <div class="col-span-2"></div>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_keitepi_5050umum_field}
        input_precision=2
        input_id="pasaran_keitepi_5050umum_field"
        input_placeholder="KEI TEPI"/>
    <Input_custom
        input_enabled={true}
        input_tipe="float"
        bind:value={pasaran_disctepi_5050umum_field}
        input_precision=2
        input_id="pasaran_disctepi_5050umum_field"
        input_placeholder="DISC TEPI(%)"/>
</div>
<div class="grid grid-cols-2 gap-2">
    <button on:click={() => {
        fetchcolok();
    }} class="{buttonLoadingfetch_class} btn-block ">Fetch</button>
    <button on:click={() => {
        save432d();
    }} class="{buttonLoading_class} btn-block ">Submit</button>
</div>