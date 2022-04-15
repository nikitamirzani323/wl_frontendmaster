<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import dayjs from "dayjs";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 
 

    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let listcurrency = [];
    export let totalrecord = 0;

    let page = "Company";
    let sData = "New";
    let sDataAdmin = "New";
    let isModal_Form_New = false
    let isModal_Form_admin = false
    let isModalLoading = false
    let isModalNotif = false
    let modal_width = "max-w-xl"
    let modal_listadmin_width = "max-w-xl"
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let buttonLoading2_class  = "btn btn-primary"
    let msg_error = "";
    let idcompany = "";
    let listAdmin = [];
    let totalrecordadmin = 0;
    let tab_listadmin = "bg-sky-600 text-white"
    let tab_listpasaran = ""
    let panel_listadmin = true

    

    let searchHome = "";
    let searchListAdmin = "";
    let filterHome = [];
    let filterListAdmin = [];
    let home_create = "";
    let home_update = "";
    let select_curr_field = "";
    let select_status_field = "";
    let admin_username_field = "";
    let admin_username_enable_field = true;
    let admin_password_field = "";
    let admin_name_field = "";
    let admin_email_field = "";
    let admin_phone_field = "";
    let admin_status_field = "";
    let admin_create = "";
    let admin_update = "";
    let admin_username_field_error = "";
    let admin_password_field_error = "";
    let admin_name_field_error = "";
    let admin_email_field_error = "";
    let admin_phone_field_error = "";
    let admin_status_field_error = "";

   
    

    

    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        home_id_field: yup
            .string()
            .required("IDCOMP is Required")
            .matches(
                /^[A-z0-9]+$/,
                "IDCOMP must Character A-Z "
            )
            .min(3, "IDCOMP must be at least 4 Character")
            .max(10, "IDCOMP must be at most 6 Character"),
        home_name_field: yup
            .string()
            .required("Company is Required")
            .matches(
                /^[A-z0-9 ]+$/,
                "Company must Character A-Z  or 1-9"
            )
            .min(4, "Company must be at least 4 Character")
            .max(70, "Company must be at most 70 Character"),
        home_nameowner_field: yup
            .string()
            .required("Owner Name is Required")
            .matches(
                /^[a-zA-z0-9 ]+$/,
                "Owner Name must Character A-Z  or 1-9"
            )
            .min(4, "Owner Name must be at least 4 Character")
            .max(70, "Owner Name must be at most 70 Character"),
        home_phoneowner_field: yup
            .string()
            .required("Owner Phone is Required")
            .min(4, "Owner Phone must be at least 4 Character")
            .max(30, "Owner Phone must be at most 30 Character"),
        home_emailowner_field: yup.string(),
        home_url_field: yup
            .string()
            .required("Url is Required")
            .min(4, "Url must be at least 4 Character")
            .max(350, "Url must be at most 350 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            home_id_field: "",
            home_name_field: "",
            home_nameowner_field: "",
            home_phoneowner_field: "",
            home_emailowner_field: "",
            home_url_field: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.home_id_field,
                values.home_name_field,
                values.home_nameowner_field,
                values.home_phoneowner_field,
                values.home_emailowner_field,
                values.home_url_field);
        },
    });
    
    async function SaveTransaksi(idcomp,namecomp,nameowner,phoneowner,emailowner,urlendpoint) {
        let flag = true;
        msg_error = "";
        if(select_curr_field == ""){
            flag = false;
            msg_error +="The Currency is required<br>";
        }
        if(select_status_field == ""){
            flag = false;
            msg_error +="The Status is required<br>";
        }
        if (flag) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompany", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "COMPANY-SAVE",
                    idcomp: idcomp,
                    idcurr: select_curr_field,
                    nmcompany: namecomp,
                    nmowner: nameowner,
                    phoneowner: phoneowner,
                    emailowner: emailowner,
                    urlendpoint: urlendpoint,
                    status: select_status_field,
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
                        $form.home_name_field = "";
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
                isModal_Form_New = false;
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    async function SaveTransaksiListAdmin() {
        let flag = false;
        msg_error = "";
       
        admin_username_field_error = "";
        admin_password_field_error = "";
        admin_name_field_error = "";
        admin_email_field_error = "";
        admin_phone_field_error = "";
        admin_status_field_error = "";
        const regexExp = /^[a-z0-9]+$/gi;
        const regexExp2 = /^[A-Za-z0-9 ]+$/gi;
        let flag_username_pattern = regexExp.test(admin_username_field)
        let flag_name_pattern = regexExp2.test(admin_name_field)
        if(admin_username_field == ""){
            flag = true
            admin_username_field_error ="Username is required"
        }
        if(!flag_username_pattern){
            flag = true
            admin_username_field_error ="Format Username must Character a-z atau 0-9"
        }
        if(sDataAdmin == "New"){
            if(admin_password_field == ""){
                flag = true
                admin_password_field_error ="Password is required"
            }
        }
        if(admin_name_field == ""){
            flag = true
            admin_name_field_error ="Name is required"
        }
        if(!flag_name_pattern){
            flag = true
            admin_name_field_error ="Format Name must Character a-z atau A-Z atau 0-9 atau spasi"
        }
        if(admin_phone_field == ""){
            flag = true
            admin_phone_field_error ="Phone is required"
        }
        if(admin_status_field == ""){
            flag = true
            admin_status_field_error ="Status is required"
        }
        if(flag == false){
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompanylistadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataAdmin,
                    page: "COMPANY-SAVE",
                    company: idcompany,
                    username: admin_username_field.toLowerCase(),
                    password: admin_password_field,
                    name: admin_name_field,
                    email: admin_email_field,
                    phone: admin_phone_field,
                    status: admin_status_field,
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
                    if(sDataAdmin == "New"){
                        clearFieldListAdmin();
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
                call_listadmin();
            }
        }
    };
    async function call_listadmin() {
        listAdmin = [];
        const res = await fetch(path_api+"api/companylistadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                totalrecordadmin = record.length;
                let company_admin_status_class = "";
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["companyadmin_status"] == "ACTIVE"){
                        company_admin_status_class = "bg-[#8BC34A] text-black"
                    }else{
                        company_admin_status_class = "bg-red-600 text-white"
                    }
                    listAdmin = [
                        ...listAdmin,
                        {
                            companyadmin_username:record[i]["companyadmin_username"],
                            companyadmin_name:record[i]["companyadmin_name"],
                            companyadmin_phone:record[i]["companyadmin_phone"],
                            companyadmin_email:record[i]["companyadmin_email"],
                            companyadmin_type:record[i]["companyadmin_type"],
                            companyadmin_status:record[i]["companyadmin_status"],
                            companyadmin_status_class:company_admin_status_class,
                            companyadmin_lastlogin:record[i]["companyadmin_lastlogin"],
                            companyadmin_lastipaddress:record[i]["companyadmin_lastipaddress"],
                            companyadmin_create:record[i]["companyadmin_create"],
                            companyadmin_update:record[i]["companyadmin_update"],
                        },
                    ];
                }
            }
        } 
    }
    

    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EntryData = (tipeentry,idcomp,nmcompany,ownername,ownerphone,owneremail,endpoint,curr,status,create,update) => {
        if(tipeentry == "Edit"){
            clearField();
            
            sData = "Edit";
            modal_width = "max-w-7xl"
            if(status == "ACTIVE"){
                status = "Y"
            }else{
                status = "N"
            }
            idcompany = idcomp;
            $form.home_id_field = idcomp;
            $form.home_name_field = nmcompany;
            $form.home_nameowner_field = ownername;
            $form.home_phoneowner_field = ownerphone;
            $form.home_emailowner_field = owneremail;
            $form.home_url_field = endpoint;
            select_curr_field = curr;
            select_status_field = status;
            home_create = create;
            home_update = update;
            call_listadmin();
        }else{
            sData = "New";
            modal_width = "max-w-2xl"
            clearField()
        }
        isModal_Form_New = true;
    };
    const NewData = () => {
        EntryData("New","","","","","","","","","","")
    };
    const EntryDataAdmin = (e,usernameadmincomp,nmadmincomp,phoneadmincomp,emailadmincomp,stausadmincomp,create,update) => {
        sDataAdmin = e;
        modal_listadmin_width = "max-w-xl"
        isModal_Form_admin = true;
        if(e == "Edit"){
            admin_username_enable_field = false
            admin_username_field = usernameadmincomp;
            admin_password_field = "";
            admin_name_field = nmadmincomp;
            admin_phone_field = phoneadmincomp;
            admin_email_field = emailadmincomp;
            admin_status_field = stausadmincomp;
            admin_create = create;
            admin_update = update;
        }else{
            admin_username_enable_field = true;
            admin_username_field = "";
            clearFieldListAdmin();
        }
    };
    
    const ChangeTabMenu = (e) => {
        switch(e){
            case "menu_listadmin":
                call_listadmin();
                tab_listadmin = "bg-sky-600 text-white"
                tab_listpasaran = ""
                panel_listadmin = true
                panel_listpasaran = false
                break;
            case "menu_listmember":
                tab_listadmin = ""
                tab_listpasaran = "bg-sky-600 text-white"
                panel_listadmin = false
                panel_listpasaran = true
                break;
        }
    }
    
    function clearField(){
        idcompany = "";
        $form.home_id_field = "";
        $form.home_name_field = "";
        $form.home_nameowner_field = "";
        $form.home_phoneowner_field = "";
        $form.home_emailowner_field = "";
        $form.home_url_field = "";
        $errors.home_id_field = "";
        $errors.home_name_field = "";
        $errors.home_nameowner_field = "";
        $errors.home_phoneowner_field = "";
        $errors.home_emailowner_field = "";
        $errors.home_url_field = "";
        select_curr_field = "";
        select_status_field = "";
    }
    function clearFieldListAdmin(){
        admin_password_field = "";
        admin_name_field = "";
        admin_email_field = "";
        admin_phone_field = "";
        admin_status_field = "";
        admin_username_field_error = "";
        admin_password_field_error = "";
        admin_name_field_error = "";
        admin_email_field_error = "";
        admin_phone_field_error = "";
        admin_status_field_error = "";
        admin_create = "";
        admin_update = "";
    }
   
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_nama
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }

        if (searchListAdmin) {
            filterListAdmin = listAdmin.filter(
                (item) =>
                    item.companyadmin_username
                        .toLowerCase()
                        .includes(searchListAdmin.toLowerCase())
            );
        } else {
            filterListAdmin = [...listAdmin];
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
            type="text" placeholder="Search by IDComp, Company" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
        <table class="table table-compact w-full ">
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NO</th>
                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">&nbsp;</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">START</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">END</th>
                    <th width="5%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">IDCOMP</th>
                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">COMPANY</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">OWNER</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">EMAIL</th>
                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">PHONE</th>
                </tr>
            </thead>
            {#if filterHome != ""}
                <tbody>
                    {#each filterHome as rec}
                    <tr>
                        <td on:click={() => {
                            EntryData("Edit",rec.home_idcompany,rec.home_name,rec.home_owner,rec.home_phone,rec.home_email,rec.home_urlendpoint,
                                rec.home_curr,rec.home_status,rec.home_create,rec.home_update);
                            }} class="text-center text-xs cursor-pointer">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                            </svg>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-center">{rec.home_no}</td>
                        <td class="text-xs lg:text-sm align-top text-center">
                            <span class="{rec.home_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.home_status}</span>
                        </td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_startjoin}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_endjoin}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_idcompany}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_name}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_owner}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_email}</td>
                        <td class="text-xs lg:text-sm align-top text-left">{rec.home_phone}</td>
                        
                    </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="12" class="text-center">
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
    modal_popup_class=" w-11/12 {modal_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        {#if sData=="New"}
            <div class="grid grid-cols-2 gap-3 mt-2 w-full">
                <div>
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={true}
                        input_required={true}
                        input_tipe="text"
                        input_text_class="uppercase"
                        input_maxlength_text="10"
                        input_invalid={$errors.home_id_field.length > 0}
                        bind:value={$form.home_id_field}
                        input_id="home_id_field"
                        input_placeholder="IDCOMP"/>
                    {#if $errors.home_id_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_id_field}</small>
                    {/if}
                </div>
                <div>
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_text_class=""
                        input_maxlength_text="70"
                        input_invalid={$errors.home_nameowner_field.length > 0}
                        bind:value={$form.home_nameowner_field}
                        input_id="home_nameowner_field"
                        input_placeholder="Owner Name"/>
                    {#if $errors.home_nameowner_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_nameowner_field}</small>
                    {/if}
                </div>
                <div>
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_text_class="uppercase"
                        input_maxlength_text="70"
                        input_invalid={$errors.home_name_field.length > 0}
                        bind:value={$form.home_name_field}
                        input_id="home_name_field"
                        input_placeholder="Company"/>
                    {#if $errors.home_name_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_name_field}</small>
                    {/if}
                </div>
                <div>
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_text_class=""
                        input_maxlength_text="70"
                        input_invalid={$errors.home_phoneowner_field.length > 0}
                        bind:value={$form.home_phoneowner_field}
                        input_id="home_phoneowner_field"
                        input_placeholder="Owner Phone"/>
                    {#if $errors.home_phoneowner_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_phoneowner_field}</small>
                    {/if}
                </div>
                <div>
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={true}
                        input_tipe="text"
                        input_maxlength_text="350"
                        input_invalid={$errors.home_url_field.length > 0}
                        bind:value={$form.home_url_field}
                        input_id="home_url_field"
                        input_placeholder="URL Endpoint"/>
                    {#if $errors.home_url_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_url_field}</small>
                    {/if}
                </div>
                <div>
                    <Input_custom
                        input_onchange="{handleChange}"
                        input_autofocus={false}
                        input_required={false}
                        input_tipe="text"
                        input_text_class=""
                        input_maxlength_text="70"
                        input_invalid={$errors.home_emailowner_field.length > 0}
                        bind:value={$form.home_emailowner_field}
                        input_id="home_emailowner_field"
                        input_placeholder="Owner Email"/>
                    {#if $errors.home_emailowner_field}
                        <small class="text-pink-600 text-[11px]">{$errors.home_emailowner_field}</small>
                    {/if}
                </div>
                <div>
                    <select
                        bind:value="{select_curr_field}" 
                        class="select select-bordered w-full focus:ring-0 focus:outline-none rounded-sm">
                        <option disabled value="" selected>--Pilih Currency--</option>
                        {#each listcurrency as rec }
                            <option value="{rec.curr_idcurr}">{rec.curr_idcurr}</option>
                        {/each}
                    </select>
                </div>
                <div>
                    <select
                        bind:value="{select_status_field}" 
                        class="select select-bordered w-full focus:ring-0 focus:outline-none rounded-sm">
                        <option disabled value="" selected>--Pilih Status--</option>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                <div class="col-span-2">
                    <button
                    on:click={() => {
                        handleSubmit();
                    }}  
                    class="{buttonLoading_class} btn-block">Submit</button>
                </div>    
            </div>
            
        {/if}
        {#if sData=="Edit"}
            <div class="flex justify-between w-full gap-2">
                <div class="w-2/3">
                    <div class="grid grid-cols-2 gap-3 mt-2 w-full">
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_enabled={false}
                                input_tipe="text"
                                input_text_class="uppercase"
                                input_maxlength_text="10"
                                input_invalid={$errors.home_id_field.length > 0}
                                bind:value={$form.home_id_field}
                                input_id="home_id_field"
                                input_placeholder="IDCOMP"/>
                            {#if $errors.home_id_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_id_field}</small>
                            {/if}
                        </div>
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_tipe="text"
                                input_text_class=""
                                input_maxlength_text="70"
                                input_invalid={$errors.home_nameowner_field.length > 0}
                                bind:value={$form.home_nameowner_field}
                                input_id="home_nameowner_field"
                                input_placeholder="Owner Name"/>
                            {#if $errors.home_nameowner_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_nameowner_field}</small>
                            {/if}
                        </div>
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_tipe="text"
                                input_text_class="uppercase"
                                input_maxlength_text="70"
                                input_invalid={$errors.home_name_field.length > 0}
                                bind:value={$form.home_name_field}
                                input_id="home_name_field"
                                input_placeholder="Company"/>
                            {#if $errors.home_name_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_name_field}</small>
                            {/if}
                        </div>
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_tipe="text"
                                input_text_class=""
                                input_maxlength_text="70"
                                input_invalid={$errors.home_phoneowner_field.length > 0}
                                bind:value={$form.home_phoneowner_field}
                                input_id="home_phoneowner_field"
                                input_placeholder="Owner Phone"/>
                            {#if $errors.home_phoneowner_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_phoneowner_field}</small>
                            {/if}
                        </div>
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={true}
                                input_tipe="text"
                                input_maxlength_text="350"
                                input_invalid={$errors.home_url_field.length > 0}
                                bind:value={$form.home_url_field}
                                input_id="home_url_field"
                                input_placeholder="URL Endpoint"/>
                            {#if $errors.home_url_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_url_field}</small>
                            {/if}
                        </div>
                        <div>
                            <Input_custom
                                input_onchange="{handleChange}"
                                input_autofocus={false}
                                input_required={false}
                                input_tipe="text"
                                input_text_class=""
                                input_maxlength_text="70"
                                input_invalid={$errors.home_emailowner_field.length > 0}
                                bind:value={$form.home_emailowner_field}
                                input_id="home_emailowner_field"
                                input_placeholder="Owner Email"/>
                            {#if $errors.home_emailowner_field}
                                <small class="text-pink-600 text-[11px]">{$errors.home_emailowner_field}</small>
                            {/if}
                        </div>
                        <div>
                            <select
                                bind:value="{select_curr_field}" 
                                class="select select-bordered w-full focus:ring-0 focus:outline-none rounded-sm">
                                <option disabled value="" selected>--Pilih Currency--</option>
                                {#each listcurrency as rec }
                                    <option value="{rec.curr_idcurr}">{rec.curr_idcurr}</option>
                                {/each}
                            </select>
                        </div>
                        <div>
                            <select
                                bind:value="{select_status_field}" 
                                class="select select-bordered w-full focus:ring-0 focus:outline-none rounded-sm">
                                <option disabled value="" selected>--Pilih Status--</option>
                                <option value="Y">ACTIVE</option>
                                <option value="N">DEACTIVE</option>
                            </select>
                        </div>
                        <div class="text-[11px] col-span-2">
                            Create : {home_create}
                            {#if home_update != ""}
                            <br>
                            Update : {home_update}
                            {/if}
                        </div>
                        <div class="col-span-2">
                            <button
                            on:click={() => {
                                handleSubmit();
                            }}  
                            class="{buttonLoading_class} btn-block">Submit</button>
                        </div>    
                    </div>
                </div>
                <div class="w-full p-2">
                    <ul class="flex justify-center items-center gap-2">
                        <li on:click={() => {
                                ChangeTabMenu("menu_listadmin");
                            }}
                            class="items-center {tab_listadmin} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Admin</li>
                        <li on:click={() => {
                                ChangeTabMenu("menu_listpasaran");
                            }}
                            class="items-center {tab_listpasaran} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Member</li>
                            <li on:click={() => {
                                ChangeTabMenu("menu_listpasaran");
                            }}
                            class="items-center {tab_listpasaran} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Transaksi</li>
                    </ul>
                    {#if panel_listadmin}
                        <h2 class="text-lg font-bold">List Admin</h2>
                        <div class="flex justify-between items-center w-full gap-1">
                            <input 
                                bind:value={searchListAdmin}
                                type="text" placeholder="Search by Username" class="input input-bordered w-full  rounded-md  focus:ring-0 focus:outline-none">
                            <button on:click={() => {
                                EntryDataAdmin("New","","","","","");
                                }}  class="btn bg-primary hover:bg-primary  rounded-md shadow-lg shadow-[#6eb5d8] border-none  ">New</button>
                        </div>
                        <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[400px] overflow-y-scroll mt-2">
                            <table class="table table-compact w-full">
                                <thead class="sticky top-0">
                                    <tr>
                                        <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">&nbsp;</th>
                                        <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">STATUS</th>
                                        <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TYPE</th>
                                        <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">LASTLOGIN</th>
                                        <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">LASTIPADDRESS</th>
                                        <th width="*" class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#if filterListAdmin != ""}
                                        {#each filterListAdmin as rec}
                                            <tr>
                                                <td class="cursor-pointer" on:click={() => {
                                                        EntryDataAdmin("Edit",rec.companyadmin_username,rec.companyadmin_name,rec.companyadmin_phone,rec.companyadmin_email,
                                                        rec.companyadmin_status,rec.companyadmin_create,rec.companyadmin_update);
                                                    }}>
                                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                                        <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                                                    </svg>
                                                </td>
                                                <td class="text-xs text-center align-top">
                                                    <span class="{rec.companyadmin_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.companyadmin_status}</span>
                                                </td>
                                                <td class="text-xs text-left align-top">{rec.companyadmin_type}</td>
                                                <td class="text-xs text-center align-top">{rec.companyadmin_lastlogin}</td>
                                                <td class="text-xs text-center align-top">{rec.companyadmin_lastipaddress}</td>
                                                <td class="text-xs text-left align-top">{rec.companyadmin_username}</td>
                                            </tr>
                                        {/each}
                                    {:else}
                                        <tr>
                                            <td colspan="5" class="text-xs">No Records</td>
                                        </tr>
                                    {/if}
                                </tbody>
                            </table>
                        </div>
                        <div class="bg-[#F7F7F7] rounded-sm h-10 p-2">
                            <table class=" w-full">
                                <tr>
                                    <td class="text-xs font-semibold text-left align-top">TOTAL RECORD</td>
                                    <td class="text-xs font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(totalrecordadmin)}</td>
                                </tr>
                            </table>
                        </div>
                    {/if}
                    
                </div>
            </div>
        {/if}
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formadmin" class="modal-toggle" bind:checked={isModal_Form_admin}>
<Modal_popup
    modal_popup_id="my-modal-formadmin"
    modal_popup_title="Admin Entry/{sDataAdmin}"
    modal_popup_class="select-none w-11/12 {modal_listadmin_width} scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-col gap-2">
            <div class="mt-2">
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_enabled={admin_username_enable_field}
                    input_tipe="text"
                    input_text_class="lowercase"
                    input_maxlength_text="40"
                    bind:value={admin_username_field}
                    input_id="admin_username_field"
                    input_placeholder="Username"/>
                {#if admin_username_field_error != ""}
                    <small class="text-pink-600 text-[11px]">{admin_username_field_error}</small>
                {/if}
            </div>
            <div>
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_tipe="password"
                    input_attr="password"
                    bind:value={admin_password_field}
                    input_id="admin_password_field"
                    input_placeholder="Password"/>
                    {#if admin_password_field_error != ""}
                        <small class="text-pink-600 text-[11px]">{admin_password_field_error}</small>
                    {/if}
            </div>
            <div>
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_enabled={true}
                    input_tipe="text"
                    input_text_class=""
                    input_maxlength_text="70"
                    bind:value={admin_name_field}
                    input_id="admin_name_field"
                    input_placeholder="Name"/>
                    {#if admin_name_field_error != ""}
                        <small class="text-pink-600 text-[11px]">{admin_name_field_error}</small>
                    {/if}
            </div>
            <div>
                <Input_custom
                    input_autofocus={false}
                    input_required={true}
                    input_enabled={true}
                    input_tipe="text"
                    input_text_class=""
                    input_maxlength_text="30"
                    bind:value={admin_phone_field}
                    input_id="admin_phone_field"
                    input_placeholder="Phone"/>
                    {#if admin_phone_field_error != ""}
                        <small class="text-pink-600 text-[11px]">{admin_phone_field_error}</small>
                    {/if}
            </div>
            <div>
                <Input_custom
                    input_autofocus={false}
                    input_required={false}
                    input_enabled={true}
                    input_tipe="text"
                    input_text_class=""
                    input_maxlength_text="150"
                    bind:value={admin_email_field}
                    input_id="admin_email_field"
                    input_placeholder="Email"/>
                    {#if admin_email_field_error != ""}
                        <small class="text-pink-600 text-[11px]">{admin_email_field_error}</small>
                    {/if}
            </div>
            <div>
                <select class="select select-bordered w-full rounded-md focus:ring-0 focus:outline-none" bind:value="{admin_status_field}">
                    <option disabled selected value="">--Pilih Status--</option>
                    <option value="ACTIVE">ACTIVE</option>
                    <option value="DEACTIVE">DEACTIVE</option>
                </select>
                {#if admin_status_field_error != ""}
                    <small class="text-pink-600 text-[11px]">{admin_status_field_error}</small>
                {/if}
            </div>
            {#if sDataAdmin == "Edit"}
            <div class="text-[11px]">
                Create : {admin_create}
                {#if admin_update != ""}
                    <br>
                    Update : {admin_update}
                {/if}
            </div>
            {/if}
            <button
                on:click={() => {
                    SaveTransaksiListAdmin();
                }}  
                class="{buttonLoading_class} btn-block">Submit</button>
        </div>
    </slot:template>
</Modal_popup>



<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />
