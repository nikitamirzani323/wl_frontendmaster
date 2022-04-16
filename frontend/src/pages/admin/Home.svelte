<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Button_custom from '../../components/button_custom.svelte'
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel_info from '../../components/Panel_info.svelte' 
    import Panel from '../../components/Panel_default.svelte' 
import { intros } from "svelte/internal";

    export let path_api = "";
    export let font_size = "";
    export let token = "";
    export let listHome = [];
    export let admin_listrule = [];
    export let totalrecord = 0;

    let page = "Admin Management";
    let sData = "New";
    let isModal_Form_New = false
    let isModalLoading = false
    let isModalNotif = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_flag = false;
    let buttonLoading_class = "btn-block";
    let msg_error = "";
    let searchHome = "";
    let filterHome = [];
    let isInput_username_enabled = true;
    let admin_create_field = ""
    let admin_update_field = ""
    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        admin_username_field: yup
            .string()
            .required("Username is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Username must Character a-z or 1-9"
            )
            .min(4, "Username must be at least 4 Character")
            .max(20, "Username must be at most 20 Character"),
        admin_password_field: yup.string(),
        admin_name_field: yup
            .string()
            .required("Nama is Required")
            .matches(
                /^[a-zA-z0-9 ]+$/,
                "Nama must Character A-Z or a-z or 1-9"
            )
            .min(4, "Nama must be at least 4 Character")
            .max(50, "Nama must be at most 50 Character"),
        admin_email_field: yup
            .string()
            .required("Email is Required")
            .min(4, "Email must be at least 4 Character")
            .max(150, "Email must be at most 150 Character"),
        admin_phone_field: yup
            .string()
            .required("Phone is Required")
            .min(4, "Phone must be at least 4 Character")
            .max(30, "Phone must be at most 30 Character"),
        admin_idrule_field: yup.string(),
        admin_status_field: yup.string().required("Status is Required"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            admin_username_field: "",
            admin_password_field: "",
            admin_name_field: "",
            admin_email_field: "",
            admin_phone_field: "",
            admin_idrule_field: "0",
            admin_status_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.admin_username_field,
                values.admin_password_field,
                values.admin_name_field,
                values.admin_email_field,
                values.admin_phone_field,
                values.admin_idrule_field,
                values.admin_status_field
            );
        },
    });
    async function SaveTransaksi(username, password, name,email,phone ,rule,status) {
        let flag = true;
        msg_error = "";
        if(status == ""){
            flag = false;
            msg_error += "The Status is required";
        }
        if (flag) {
            buttonLoading_flag = true;
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "ADMIN-SAVE",
                    username: username.toLowerCase(),
                    password: password,
                    name: name,
                    email: email,
                    phone: phone,
                    idrule: parseInt(rule),
                    status: status,
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
                        $form.admin_username_field = "";
                        $form.admin_password_field = "";
                        $form.admin_name_field = "";
                        $form.admin_idrule_field = "";
                        $form.admin_status_field = "";
                    }
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_flag = false;
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
    
    const EntryData = (tipeentry,username,name,email,phone,rule,status,create,update) => {
        if(tipeentry == "Edit"){
            sData = "Edit"
            isInput_username_enabled = false;
            $form.admin_username_field = username;
            $form.admin_password_field = "";
            $form.admin_name_field = name;
            $form.admin_email_field = email;
            $form.admin_phone_field = phone;
            $form.admin_idrule_field = rule;
            $form.admin_status_field = status;
            $errors.admin_username_field = "";
            $errors.admin_password_field = "";
            $errors.admin_name_field = "";
            $errors.admin_email_field = "";
            $errors.admin_phone_field = "";
            $errors.admin_idrule_field = "";
            $errors.admin_status_field = "";
            admin_create_field = create;
            admin_update_field = update;
            isModal_Form_New = true;
        }else{
            sData = "New"
            isInput_username_enabled = true;
            clearField()
            isModal_Form_New = true;
        }
        
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        EntryData("New","","","","","","")
        
    };
    
    function clearField(){
        $form.admin_username_field = "";
        $form.admin_password_field = "";
        $form.admin_name_field = "";
        $form.admin_email_field = "";
        $form.admin_phone_field = "";
        $form.admin_idrule_field = "0";
        $form.admin_status_field = "";
        $errors.admin_username_field = "";
        $errors.admin_password_field = "";
        $errors.admin_name_field = "";
        $errors.admin_email_field = "";
        $errors.admin_phone_field = "";
        $errors.admin_idrule_field = "";
        $errors.admin_status_field = "";
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
            type="text" placeholder="Search by Rule, Username, Nama" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center"></th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">NO</th>
                    <th width="1%" class="bg-[#475289] {font_size} text-white text-center">STATUS</th>
                    <th width="7%" class="bg-[#475289] {font_size} text-white text-center">IPADDRESS</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-center">LAST LOGIN</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-center">JOIN DATE</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-left">RULE</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-left">USERNAME</th>
                    <th width="*" class="bg-[#475289] {font_size} text-white text-left">NAMA</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-left">EMAIL</th>
                    <th width="15%" class="bg-[#475289] {font_size} text-white text-left">PHONE</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EntryData("Edit",
                                rec.home_username,rec.home_nama,rec.home_email,rec.home_phone,
                                rec.home_idrule,rec.home_status,rec.home_create,rec.home_update);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="{font_size} align-top text-center">{rec.home_no}</td>
                        <td class="{font_size} align-top text-center"><span class="{rec.home_statusclass} text-center rounded-md p-1 px-2">{rec.home_status}</span></td>
                        <td class="{font_size} align-top text-center">{rec.home_lastipaddres}</td>
                        <td class="{font_size} align-top text-center">{rec.home_lastlogin}</td>
                        <td class="{font_size} align-top text-center">{rec.home_joindate}</td>
                        <td class="{font_size} align-top text-left">{rec.home_rule}</td>
                        <td class="{font_size} align-top text-left">{rec.home_username}</td>
                        <td class="{font_size} align-top text-left">{rec.home_nama}</td>
                        <td class="{font_size} align-top text-left">{rec.home_email}</td>
                        <td class="{font_size} align-top text-left">{rec.home_phone}</td>
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
        <div class="grid grid-cols-2 gap-2 mt-2 ">
            <div class="">
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="20"
                    input_text_class="lowercase"
                    input_invalid={$errors.admin_username_field.length > 0}
                    bind:value={$form.admin_username_field}
                    input_id="admin_username_field"
                    input_enabled={isInput_username_enabled}
                    input_placeholder="Username"/>
                {#if $errors.admin_username_field}
                    <small class="text-pink-600 text-[11px]">{$errors.admin_username_field}</small>
                {/if}
            </div>
            <div class="">
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="50"
                    input_invalid={$errors.admin_name_field.length > 0}
                    bind:value={$form.admin_name_field}
                    input_id="admin_name_field"
                    input_placeholder="Nama"/>
                {#if $errors.admin_name_field}
                    <small class="text-pink-600 text-[11px]">{$errors.admin_name_field}</small>
                {/if}
            </div>
            <div class="">
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="password"
                    input_attr="password"
                    input_invalid={$errors.admin_password_field.length > 0}
                    bind:value={$form.admin_password_field}
                    input_id="admin_password_field"
                    input_placeholder="Password"/>
                {#if $errors.admin_password_field}
                    <small class="text-pink-600 text-[11px]">{$errors.admin_password_field}</small>
                {/if}
            </div>
            <div class="">
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="150"
                    input_invalid={$errors.admin_email_field.length > 0}
                    bind:value={$form.admin_email_field}
                    input_id="admin_email_field"
                    input_placeholder="Email"/>
                {#if $errors.admin_email_field}
                    <small class="text-pink-600 text-[11px]">{$errors.admin_email_field}</small>
                {/if}
            </div>
            <div class="">
                <select
                    on:change="{handleChange}"
                    bind:value={$form.admin_idrule_field}
                    invalid={$errors.admin_idrule_field.length > 0} 
                    class="w-full rounded px-3 text-sm lg:text-md  border border-gray-300 focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none">
                    <option disabled selected value="0">--Pilih Admin Rule--</option>
                    {#each admin_listrule as rec}
                    <option value="{rec.adminrule_idrule}">{rec.adminrule_nmrule}</option>
                    {/each}
                </select>
                {#if $errors.admin_idrule_field}
                    <small class="text-pink-600 text-[11px]">{$errors.admin_idrule_field}</small>
                {/if}
            </div>
            <div class="">
                <Input_custom
                    input_onchange="{handleChange}"
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="text"
                    input_maxlength_text="30"
                    input_invalid={$errors.admin_phone_field.length > 0}
                    bind:value={$form.admin_phone_field}
                    input_id="admin_phone_field"
                    input_placeholder="Phone"/>
                {#if $errors.admin_phone_field}
                    <small class="text-pink-600 text-[11px]">{$errors.admin_phone_field}</small>
                {/if}
            </div>
            <div class="">
                <select
                    on:change="{handleChange}"
                    bind:value={$form.admin_status_field}
                    invalid={$errors.admin_status_field.length > 0} 
                    class="w-full text-sm lg:text-md rounded px-3  border border-gray-300 focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none">
                    <option disabled selected value="">--Pilih Status--</option>
                    <option value="ACTIVE">ACTIVE</option>
                    <option value="BANNED">BANNED</option>
                </select>
                {#if $errors.admin_status_field}
                    <small class="text-pink-600 text-[11px]">{$errors.admin_status_field}</small>
                {/if}
            </div>
        </div>
        {#if sData == "Edit"}
            <Panel_info panel_body_class="mt-2">
                <slot:template slot="panel_body">
                    <table>
                        <tr>
                            <td>Create</td>
                            <td>:</td>
                            <td>{admin_create_field}</td>
                        </tr>
                        {#if admin_update_field != ""}
                        <tr>
                            <td>Modified</td>
                            <td>:</td>
                            <td>{admin_update_field}</td>
                        </tr>
                        {/if}
                    </table>
                </slot:template>
            </Panel_info>
        {/if}
        <div class="col-span-2 mt-2">
            <Button_custom 
                on:click={() => {
                    handleSubmit();
                }}
                button_disable={buttonLoading_flag}
                button_class="btn-block"
                button_disable_class="{buttonLoading_class}"
                button_title="Submit" />
        </div>
    </slot:template>
</Modal_popup>



<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />



