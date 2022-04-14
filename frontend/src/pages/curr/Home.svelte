<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 

    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let admin_listrule = [];
    export let totalrecord = 0;

    let page = "Currency";
    let sData = "New";
    let isModal_Form_New = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let msg_error = "";
    let searchHome = "";
    let filterHome = [];
    let isInput_idcurr_enabled = true;
    let curr_create_field = ""
    let curr_update_field = ""
    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        curr_id_field: yup
            .string()
            .required("Username is Required")
            .matches(
                /^[a-zA-z]+$/,
                "Username must Character a-z or 1-9"
            )
            .min(2, "Username must be at least 4 Character")
            .max(4, "Username must be at most 4 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            curr_id_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(values.curr_id_field);
        },
    });
    async function SaveTransaksi(idcurr) {
        let flag = true;
        msg_error = "";
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecurrency", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "ADMIN-SAVE",
                    idcurr: idcurr.toUpperCase(),
                    name: idcurr.toUpperCase(),
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
                    if(sData == "New"){
                        $form.curr_id_field = "";
                    }
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                RefreshHalaman();
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    
    const EntryData = (tipeentry,idcurr,create,update) => {
        if(tipeentry == "Edit"){
            sData = "Edit"
            isInput_idcurr_enabled = false;
            $form.curr_id_field = idcurr;
            curr_create_field = create;
            curr_update_field = update;
            isModal_Form_New = true;
        }else{
            sData = "New"
            isInput_idcurr_enabled = true;
            clearField()
            isModal_Form_New = true;
        }
        
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        EntryData("New","","","")
        
    };
    
    function clearField(){
        $form.curr_id_field = "";
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.admin_username
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) ||
                    item.admin_nama
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.admin_rule
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />

<Panel
    on:handleNewForm={NewData}
    on:handleRefresh={RefreshHalaman}
    panel_button_new={true}
    panel_button_refresh={true}
    panel_page="{page}"
    panel_total="{totalrecord}">
    <slot:template slot="panel_search">
        <div class="absolute inset-y-0 left-0 flex items-center pl-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
        </div>
        <input 
            bind:value={searchHome}
            type="text" placeholder="Search by Nama" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NO</th>
                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">CURRENCY</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EntryData("Edit",rec.home_idcurr,rec.home_create,rec.home_update);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.home_no}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_idcurr}</td>
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="10" class="text-center">
                            <progress class="self-start progress progress-primary w-56"></progress>
                        </td>
                    </tr>
                </tbody>
            {/if}
        </table>
    </slot:template>
</Panel>


<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<Modal_popup
    modal_popup_id="my-modal-formnew"
    modal_popup_title="Entry/{sData}"
    modal_popup_class="select-none max-w-full lg:max-w-xl overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2 ">
            <div class="mt-2">
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="4"
                    input_text_class="uppercase"
                    input_invalid={$errors.curr_id_field.length > 0}
                    bind:value={$form.curr_id_field}
                    input_id="admin_username_field"
                    input_enabled={isInput_idcurr_enabled}
                    input_placeholder="Currency"/>
                {#if $errors.curr_id_field}
                    <small class="text-pink-600 text-[11px]">{$errors.curr_id_field}</small>
                {/if}
            </div>
           
            {#if sData == "Edit"}
            <div class="text-[11px]">
                Create : {curr_create_field} <br>
                Update : {curr_update_field}
            </div>
            {/if}
        </div>
        {#if sData == "New"}
        <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
            <button
                on:click={() => {
                    handleSubmit();
                }}  
                class="{buttonLoading_class}">Submit</button>
        </div>
        {/if}
    </slot:template>
</Modal_popup>



<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



