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
    import Panel_432D from './432D.svelte'
    import Panel_cbebas from './colok_bebas.svelte'
    import Panel_cmacau from './colok_macau.svelte'
    import Panel_cnaga from './colok_naga.svelte'
    import Panel_cjitu from './colok_jitu.svelte'
    import Panel_5050umum from './5050_umum.svelte'
    import Panel_5050special from './5050_special.svelte'
    import Panel_5050kombinasi from './5050_kombinasi.svelte'
    import Panel_kombinasi from './kombinasi.svelte'
    import Panel_dasar from './dasar.svelte'
    import Panel_shio from './shio.svelte'

    export let path_api = "";
    export let token = "";
    export let master = "";
    export let listHome = [];
    export let listcurrency = [];
    export let totalrecord = 0;

    let page = "Company";
    let sData = "New";
    let sDataAdmin = "New";
    let isModal_Form_New = false
    let isModal_Form_admin = false
    let isModal_Form_pasaran = false
    let isModal_Form_confpasaran = false
    let isModalLoading = false
    let isModalNotif = false
    let modal_width = "max-w-xl"
    let modal_listadmin_width = "max-w-xl"
    let modal_listpasaran_width = "max-w-xl"
    let modal_confpasaran_width = "max-w-xl"
    let modalpasaran_width = "max-w-xl"
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let buttonLoading2_class  = "btn btn-primary"
    let msg_error = "";
    let home_status_field = "";
    let home_create_field = "";
    let home_update_field = "";
    let idcompany = "";
    let listAdmin = [];
    let listPasaran = [];
    let listPasaranonline = [];
    let totalrecordadmin = 0;
    let totalpasaran = 0;
    let totalpasaran_class = "";
    let tab_listadmin = "bg-sky-600 text-white"
    let tab_listpasaran = ""
    let panel_listadmin = true
    let panel_listpasaran = false

    let tab_listpasaran_limit = "bg-sky-600 text-white"
    let tab_listpasaran_online = ""
    let tab_listpasaran_configure = ""
    let tab_listpasaran_listkeluaran= ""
    let panel_listpasaran_limit = true
    let panel_listpasaran_online = false
    let panel_listpasaran_configure = false
    let panel_listpasaran_listkeluaran = false

    let searchHome = "";
    let searchListAdmin = "";
    let filterHome = [];
    let filterListAdmin = [];
    let permainan = ""; 
    let home_create = "";
    let home_update = "";
    let select_curr_field = "";
    let select_status_field = "";
    let admin_username_field = "";
    let admin_username_enable_field = true;
    let admin_password_field = "";
    let admin_name_field = "";
    let admin_status_field = "";
    let admin_username_field_error = "";
    let admin_password_field_error = "";
    let admin_name_field_error = "";
    let admin_status_field_error = "";

    let companypasaran_id = "";
    let pasaran_id_field = "";
    let pasaran_nmpasarantogel_field = "";
    let pasaran_urlpasaran_field = "";
    let pasaran_pasarandiundi_field = "";
    let pasaran_jamtutup_field = "";
    let pasaran_jamjadwal_field = "";
    let pasaran_jamopen_field = "";
    let pasaran_status_field = "";
    let select_pasaranonline = "";

    let pasaran_limitline4d_field = 0;
    let pasaran_limitline3d_field = 0;
    let pasaran_limitline3dd_field = 0;
    let pasaran_limitline2d_field = 0;
    let pasaran_limitline2dd_field = 0;
    let pasaran_limitline2dt_field = 0;
    let pasaran_bbfs_field = 0;
    

    let panel_432D = false
    let panel_cbebas = false
    let panel_cmacau = false
    let panel_cnaga = false
    let panel_cjitu = false
    let panel_5050umum = false
    let panel_5050special = false
    let panel_5050kombinasi = false
    let panel_macaukombinasi = false
    let panel_dasar = false
    let panel_shio = false

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
    
    async function call_listadmin() {
        listAdmin = [];
        const res = await fetch(path_api+"api/companylistadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                sData: "New",
                page: "COMPANY_HOME",
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
                    if(record[i]["company_admin_status"] == "ACTIVE"){
                        company_admin_status_class = "bg-[#8BC34A] text-black"
                    }else{
                        company_admin_status_class = "bg-red-600 text-white"
                    }
                    listAdmin = [
                        ...listAdmin,
                        {
                            company_admin_username:record[i]["company_admin_username"],
                            company_admin_typeadmin:record[i]["company_admin_typeadmin"],
                            company_admin_nama: record[i]["company_admin_nama"],
                            company_admin_status:record[i]["company_admin_status"],
                            company_admin_status_class:company_admin_status_class,
                            company_admin_lastlogin:record[i]["company_admin_lastlogin"],
                            company_admin_lastipaddres:record[i]["company_admin_lastipaddres"],
                            company_admin_create:record[i]["company_admin_create"],
                            company_admin_update:record[i]["company_admin_update"],
                        },
                    ];
                }
            }
        } 
    }
    async function call_listpasaran() {
        listPasaran = [];
        const res = await fetch(path_api+"api/companylistpasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                sData: "New",
                page: "COMPANY_HOME",
                company: idcompany,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                totalrecordadmin = record.length;
                totalpasaran = 0;
                for (var i = 0; i < record.length; i++) {
                    let selisihwinlose_class = ""
                    let status_class = ""
                    let statusactive_class = ""
                    if (record[i]["company_pasaran_winlose"] > 0) {
                        selisihwinlose_class = "text-blue-700";
                    } else {
                        selisihwinlose_class = "text-red-500";
                    }
                    if(record[i]["company_pasaran_status"] == "ONLINE"){
                        status_class = "bg-[#FFEB3B] text-black"
                    }else{
                        status_class = "bg-red-600 text-white"
                    }
                    if(record[i]["company_pasaran_statuspasaranactive"] == "Y"){
                        statusactive_class = "bg-[#8BC34A] text-black"
                    }else{
                        statusactive_class = "bg-red-600 text-white"
                    }
                    totalpasaran = totalpasaran + parseInt(record[i]["company_pasaran_winlose"])
                    listPasaran = [
                        ...listPasaran,
                        {
                            company_pasaran_idcomppasaran:record[i]["company_pasaran_idcomppasaran"],
                            company_pasaran_idpasarantogel:record[i]["company_pasaran_idpasarantogel"],
                            company_pasaran_nmpasarantogel:record[i]["company_pasaran_nmpasarantogel"],
                            company_pasaran_periode:record[i]["company_pasaran_periode"],
                            company_pasaran_winlose:record[i]["company_pasaran_winlose"],
                            company_pasaran_csswinlose:selisihwinlose_class,
                            company_pasaran_displaypasaran:record[i]["company_pasaran_displaypasaran"],
                            company_pasaran_status:record[i]["company_pasaran_status"],
                            company_pasaran_statuscss:status_class,
                            company_pasaran_statuspasaranactive:record[i]["company_pasaran_statuspasaranactive"],
                            company_pasaran_statuspasaranactivecss:statusactive_class,
                        },
                    ];
                }
                totalpasaran_class = "text-red-600 font-semibold"
                if(totalpasaran > 0){
                    totalpasaran_class = "text-blue-700 font-semibold"
                }
            }
        }
    }
    async function call_listpasaranconf() {
        const res = await fetch(path_api+"api/companypasaranconf", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                sData: "New",
                page: "COMPANY_HOME",
                company: idcompany,
                companypasaran_id: companypasaran_id,
            }),
        });
        const json = await res.json();
        let record = json.record;
        if (json.status === 400) {
        } else {
            for (let i = 0; i < record.length; i++) {
                let jamtutup = dayjs().format("DD MMM YYYY ")+record[i]["pasaran_jamtutup"];
                let jamjadwal = dayjs().format("DD MMM YYYY ")+record[i]["pasaran_jamjadwal"];
                let jamopen = dayjs().format("DD MMM YYYY ")+record[i]["pasaran_jamopen"];

                pasaran_urlpasaran_field = record[i]["pasaran_url"];
                pasaran_pasarandiundi_field = record[i]["pasaran_diundi"];
                pasaran_jamtutup_field = dayjs(jamtutup).format("HH:mm");
                pasaran_jamjadwal_field = dayjs(jamjadwal).format("HH:mm");
                pasaran_jamopen_field = dayjs(jamopen).format("HH:mm");
                pasaran_status_field = record[i]["pasaran_statusactive"];
                pasaran_bbfs_field = record[i]["bbfs"];
                pasaran_limitline4d_field = record[i]["limitline_4d"];
                pasaran_limitline3d_field = record[i]["limitline_3d"];
                pasaran_limitline3dd_field = record[i]["limitline_3dd"];
                pasaran_limitline2d_field = record[i]["limitline_2d"];
                pasaran_limitline2dd_field = record[i]["limitline_2dd"];
                pasaran_limitline2dt_field = record[i]["limitline_2dt"];
                pasaran_minbet_432d_field = record[i]["minbet_432d"];
                pasaran_maxbet4d_432d_field = record[i]["maxbet4d_432d"];
                pasaran_maxbet3d_432d_field = record[i]["maxbet3d_432d"];
                pasaran_maxbet3dd_432d_field = record[i]["maxbet3dd_432d"];
                pasaran_maxbet2d_432d_field = record[i]["maxbet2d_432d"];
                pasaran_maxbet2dd_432d_field = record[i]["maxbet2dd_432d"];
                pasaran_maxbet2dt_432d_field = record[i]["maxbet2dt_432d"];
                pasaran_limitotal4d_432d_field = record[i]["limitotal4d_432d"];
                pasaran_limitotal3d_432d_field = record[i]["limitotal3d_432d"];
                pasaran_limitotal3dd_432d_field = record[i]["limitotal3dd_432d"];
                pasaran_limitotal2d_432d_field = record[i]["limitotal2d_432d"];
                pasaran_limitotal2dd_432d_field = record[i]["limitotal2dd_432d"];
                pasaran_limitotal2dt_432d_field = record[i]["limitotal2dt_432d"];
                pasaran_limitglobal4d_432d_field = record[i]["limitglobal4d_432d"];
                pasaran_limitglobal3d_432d_field = record[i]["limitglobal3d_432d"];
                pasaran_limitglobal3dd_432d_field = record[i]["limitglobal3dd_432d"];
                pasaran_limitglobal2d_432d_field = record[i]["limitglobal2d_432d"];
                pasaran_limitglobal2dd_432d_field = record[i]["limitglobal2dd_432d"];
                pasaran_limitglobal2dt_432d_field = record[i]["limitglobal2dt_432d"];
                pasaran_disc4d_432d_field = (record[i]["disc4d_432d"] * 100).toFixed(2);
                pasaran_disc3d_432d_field = (record[i]["disc3d_432d"] * 100).toFixed(2);
                pasaran_disc3dd_432d_field = (record[i]["disc3dd_432d"] * 100).toFixed(2);
                pasaran_disc2d_432d_field = (record[i]["disc2d_432d"] * 100).toFixed(2);
                pasaran_disc2dd_432d_field = (record[i]["disc2dd_432d"] * 100).toFixed(2);
                pasaran_disc2dt_432d_field = (record[i]["disc2dt_432d"] * 100).toFixed(2);
                pasaran_win4d_432d_field = record[i]["win4d_432d"];
                pasaran_win3d_432d_field = record[i]["win3d_432d"];
                pasaran_win3dd_432d_field = record[i]["win3dd_432d"];
                pasaran_win2d_432d_field = record[i]["win2d_432d"];
                pasaran_win2dd_432d_field = record[i]["win2dd_432d"];
                pasaran_win2dt_432d_field = record[i]["win2dt_432d"];
                pasaran_win4d_nodisc_432d_field = record[i]["win4dnodisc_432d"];
                pasaran_win3d_nodisc_432d_field = record[i]["win3dnodisc_432d"];
                pasaran_win3dd_nodisc_432d_field = record[i]["win3ddnodisc_432d"];
                pasaran_win2d_nodisc_432d_field = record[i]["win2dnodisc_432d"];
                pasaran_win2dd_nodisc_432d_field = record[i]["win2ddnodisc_432d"];
                pasaran_win2dt_nodisc_432d_field = record[i]["win2dtnodisc_432d"];
                pasaran_win4d_bb_kena_432d_field = record[i]["win4dbb_kena_432d"];
                pasaran_win3d_bb_kena_432d_field = record[i]["win3dbb_kena_432d"];
                pasaran_win3dd_bb_kena_432d_field = record[i]["win3ddbb_kena_432d"];
                pasaran_win2d_bb_kena_432d_field = record[i]["win2dbb_kena_432d"];
                pasaran_win2dd_bb_kena_432d_field = record[i]["win2ddbb_kena_432d"];
                pasaran_win2dt_bb_kena_432d_field = record[i]["win2dtbb_kena_432d"];
                pasaran_win4d_bb_432d_field = record[i]["win4dbb_432d"];
                pasaran_win3d_bb_432d_field = record[i]["win3dbb_432d"];
                pasaran_win3dd_bb_432d_field = record[i]["win3ddbb_432d"];
                pasaran_win2d_bb_432d_field = record[i]["win2dbb_432d"];
                pasaran_win2dd_bb_432d_field = record[i]["win2ddbb_432d"];
                pasaran_win2dt_bb_432d_field = record[i]["win2dtbb_432d"];
                pasaran_minbet_cbebas_field = record[i]["minbet_cbebas"];
                pasaran_maxbet_cbebas_field = record[i]["maxbet_cbebas"];
                pasaran_limitotal_cbebas_field = record[i]["limittotal_cbebas"];
                pasaran_limitglobal_cbebas_field = record[i]["limitglobal_cbebas"];
                pasaran_win_cbebas_field = record[i]["win_cbebas"].toFixed(2);
                pasaran_disc_cbebas_field = (record[i]["disc_cbebas"] * 100).toFixed(2);
                pasaran_minbet_cmacau_field = record[i]["minbet_cmacau"];
                pasaran_maxbet_cmacau_field = record[i]["maxbet_cmacau"];
                pasaran_limitotal_cmacau_field = record[i]["limitotal_cmacau"];
                pasaran_limitglobal_cmacau_field = record[i]["limitglobal_cmacau"];
                pasaran_win2_cmacau_field = record[i]["win2d_cmacau"].toFixed(2);
                pasaran_win3_cmacau_field = record[i]["win3d_cmacau"].toFixed(2);
                pasaran_win4_cmacau_field = record[i]["win4d_cmacau"].toFixed(2);
                pasaran_disc_cmacau_field = (record[i]["disc_cmacau"] * 100).toFixed(2);
                pasaran_minbet_cnaga_field = record[i]["minbet_cnaga"];
                pasaran_maxbet_cnaga_field = record[i]["maxbet_cnaga"];
                pasaran_win3_cnaga_field = record[i]["win3_cnaga"].toFixed(2);
                pasaran_win4_cnaga_field = record[i]["win4_cnaga"].toFixed(2);
                pasaran_disc_cnaga_field = (record[i]["disc_cnaga"] * 100).toFixed(2);
                pasaran_limitglobal_cnaga_field = record[i]["limitglobal_cnaga"];
                pasaran_limittotal_cnaga_field = record[i]["limittotal_cnaga"];
                pasaran_minbet_cjitu_field = record[i]["minbet_cjitu"];
                pasaran_maxbet_cjitu_field = record[i]["maxbet_cjitu"];
                pasaran_winas_cjitu_field = record[i]["winas_cjitu"].toFixed(2);
                pasaran_winkop_cjitu_field = record[i]["winkop_cjitu"].toFixed(2);
                pasaran_winkepala_cjitu_field = record[i]["winkepala_cjitu"].toFixed(2);
                pasaran_winekor_cjitu_field = record[i]["winekor_cjitu"].toFixed(2);
                pasaran_desc_cjitu_field = (record[i]["desc_cjitu"] * 100).toFixed(2);
                pasaran_limitglobal_cjitu_field = record[i]["limitglobal_cjitu"];
                pasaran_limittotal_cjitu_field = record[i]["limittotal_cjitu"];
                pasaran_minbet_5050umum_field = record[i]["minbet_5050umum"];
                pasaran_maxbet_5050umum_field = record[i]["maxbet_5050umum"];
                pasaran_keibesar_5050umum_field = (record[i]["keibesar_5050umum"] * 100).toFixed(2);
                pasaran_keikecil_5050umum_field = (record[i]["keikecil_5050umum"] * 100).toFixed(2);
                pasaran_keigenap_5050umum_field = (record[i]["keigenap_5050umum"] * 100).toFixed(2);
                pasaran_keiganjil_5050umum_field = (record[i]["keiganjil_5050umum"] * 100).toFixed(2);
                pasaran_keitengah_5050umum_field = (record[i]["keitengah_5050umum"] * 100).toFixed(2);
                pasaran_keitepi_5050umum_field = (record[i]["keitepi_5050umum"] * 100).toFixed(2);
                pasaran_discbesar_5050umum_field = (record[i]["discbesar_5050umum"] * 100).toFixed(2);
                pasaran_disckecil_5050umum_field = (record[i]["disckecil_5050umum"] * 100).toFixed(2);
                pasaran_discgenap_5050umum_field = (record[i]["discgenap_5050umum"] * 100).toFixed(2);
                pasaran_discganjil_5050umum_field = (record[i]["discganjil_5050umum"] * 100).toFixed(2);
                pasaran_disctengah_5050umum_field = (record[i]["disctengah_5050umum"] * 100).toFixed(2);
                pasaran_disctepi_5050umum_field = (record[i]["disctepi_5050umum"] * 100).toFixed(2);
                pasaran_limitglobal_5050umum_field = record[i]["limitglobal_5050umum"];
                pasaran_limittotal_5050umum_field = record[i]["limittotal_5050umum"];
                
                pasaran_minbet_5050special_field = record[i]["minbet_5050special"];
                pasaran_maxbet_5050special_field = record[i]["maxbet_5050special"];
                pasaran_keiasganjil_5050special_field = (record[i]["keiasganjil_5050special"] * 100).toFixed(2);
                pasaran_keiasgenap_5050special_field = (record[i]["keiasgenap_5050special"] * 100).toFixed(2);
                pasaran_keiasbesar_5050special_field = (record[i]["keiasbesar_5050special"] * 100).toFixed(2);
                pasaran_keiaskecil_5050special_field = (record[i]["keiaskecil_5050special"] * 100).toFixed(2);
                pasaran_keikopganjil_5050special_field = (record[i]["keikopganjil_5050special"] * 100).toFixed(2);
                pasaran_keikopgenap_5050special_field = (record[i]["keikopgenap_5050special"] * 100).toFixed(2);
                pasaran_keikopbesar_5050special_field = (record[i]["keikopbesar_5050special"] * 100).toFixed(2);
                pasaran_keikopkecil_5050special_field = (record[i]["keikopkecil_5050special"] * 100).toFixed(2);
                pasaran_keikepalaganjil_5050special_field = (record[i]["keikepalaganjil_5050special"] * 100).toFixed(2);
                pasaran_keikepalagenap_5050special_field = (record[i]["keikepalagenap_5050special"] * 100).toFixed(2);
                pasaran_keikepalabesar_5050special_field = (record[i]["keikepalabesar_5050special"] * 100).toFixed(2);
                pasaran_keikepalakecil_5050special_field = (record[i]["keikepalakecil_5050special"] * 100).toFixed(2);
                pasaran_keiekorganjil_5050special_field = (record[i]["keiekorganjil_5050special"] * 100).toFixed(2);
                pasaran_keiekorgenap_5050special_field = (record[i]["keiekorgenap_5050special"] * 100).toFixed(2);
                pasaran_keiekorbesar_5050special_field = (record[i]["keiekorbesar_5050special"] * 100).toFixed(2);
                pasaran_keiekorkecil_5050special_field = (record[i]["keiekorkecil_5050special"] * 100).toFixed(2);
                pasaran_discasganjil_5050special_field = (record[i]["discasganjil_5050special"] * 100).toFixed(2);
                pasaran_discasgenap_5050special_field = (record[i]["discasgenap_5050special"] * 100).toFixed(2);
                pasaran_discasbesar_5050special_field = (record[i]["discasbesar_5050special"] * 100).toFixed(2);
                pasaran_discaskecil_5050special_field = (record[i]["discaskecil_5050special"] * 100).toFixed(2);
                pasaran_disckopganjil_5050special_field = (record[i]["disckopganjil_5050special"] * 100).toFixed(2);
                pasaran_disckopgenap_5050special_field = (record[i]["disckopgenap_5050special"] * 100).toFixed(2);
                pasaran_disckopbesar_5050special_field = (record[i]["disckopbesar_5050special"] * 100).toFixed(2);
                pasaran_disckopkecil_5050special_field = (record[i]["disckopkecil_5050special"] * 100).toFixed(2);
                pasaran_disckepalaganjil_5050special_field = (record[i]["disckepalaganjil_5050special"] * 100).toFixed(2);
                pasaran_disckepalagenap_5050special_field = (record[i]["disckepalagenap_5050special"] * 100).toFixed(2);
                pasaran_disckepalabesar_5050special_field = (record[i]["disckepalabesar_5050special"] * 100).toFixed(2);
                pasaran_disckepalakecil_5050special_field = (record[i]["disckepalakecil_5050special"] * 100).toFixed(2);
                pasaran_discekorganjil_5050special_field = (record[i]["discekorganjil_5050special"] * 100).toFixed(2);
                pasaran_discekorgenap_5050special_field = (record[i]["discekorgenap_5050special"] * 100).toFixed(2);
                pasaran_discekorbesar_5050special_field = (record[i]["discekorbesar_5050special"] * 100).toFixed(2);
                pasaran_discekorkecil_5050special_field = (record[i]["discekorkecil_5050special"] * 100).toFixed(2);
                pasaran_limitglobal_5050special_field = record[i]["limitglobal_5050special"];
                pasaran_limittotal_5050special_field = record[i]["limittotal_5050special"];
                pasaran_minbet_5050kombinasi_field = record[i]["minbet_5050kombinasi"];
                pasaran_maxbet_5050kombinasi_field = record[i]["maxbet_5050kombinasi"];
                pasaran_belakangkeimono_5050kombinasi_field = (record[i]["belakangkeimono_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangkeistereo_5050kombinasi_field = (record[i]["belakangkeistereo_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangkeikembang_5050kombinasi_field = (record[i]["belakangkeikembang_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangkeikempis_5050kombinasi_field = (record[i]["belakangkeikempis_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangkeikembar_5050kombinasi_field = (record[i]["belakangkeikembar_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahkeimono_5050kombinasi_field = (record[i]["tengahkeimono_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahkeistereo_5050kombinasi_field = (record[i]["tengahkeistereo_5050kombinasi"] * 100).toFixed(2)
                pasaran_tengahkeikembang_5050kombinasi_field = (record[i]["tengahkeikembang_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahkeikempis_5050kombinasi_field = (record[i]["tengahkeikempis_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahkeikembar_5050kombinasi_field = (record[i]["tengahkeikembar_5050kombinasi"] * 100).toFixed(2);
                pasaran_depankeimono_5050kombinasi_field = (record[i]["depankeimono_5050kombinasi"] * 100).toFixed(2);
                pasaran_depankeistereo_5050kombinasi_field = (record[i]["depankeistereo_5050kombinasi"] * 100).toFixed(2);
                pasaran_depankeikembang_5050kombinasi_field = (record[i]["depankeikembang_5050kombinasi"] * 100).toFixed(2);
                pasaran_depankeikempis_5050kombinasi_field = (record[i]["depankeikempis_5050kombinasi"] * 100).toFixed(2);
                pasaran_depankeikembar_5050kombinasi_field = (record[i]["depankeikembar_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangdiscmono_5050kombinasi_field = (record[i]["belakangdiscmono_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangdiscstereo_5050kombinasi_field = (record[i]["belakangdiscstereo_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangdisckembang_5050kombinasi_field = (record[i]["belakangdisckembang_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangdisckempis_5050kombinasi_field = (record[i]["belakangdisckempis_5050kombinasi"] * 100).toFixed(2);
                pasaran_belakangdisckembar_5050kombinasi_field = (record[i]["belakangdisckembar_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahdiscmono_5050kombinasi_field = (record[i]["tengahdiscmono_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahdiscstereo_5050kombinasi_field = (record[i]["tengahdiscstereo_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahdisckembang_5050kombinasi_field = (record[i]["tengahdisckembang_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahdisckempis_5050kombinasi_field = (record[i]["tengahdisckempis_5050kombinasi"] * 100).toFixed(2);
                pasaran_tengahdisckembar_5050kombinasi_field = (record[i]["tengahdisckembar_5050kombinasi"] * 100).toFixed(2);
                pasaran_depandiscmono_5050kombinasi_field = (record[i]["depandiscmono_5050kombinasi"] * 100).toFixed(2);
                pasaran_depandiscstereo_5050kombinasi_field = (record[i]["depandiscstereo_5050kombinasi"] * 100).toFixed(2);
                pasaran_depandisckembang_5050kombinasi_field = (record[i]["depandisckembang_5050kombinasi"] * 100).toFixed(2);
                pasaran_depandisckempis_5050kombinasi_field = (record[i]["depandisckempis_5050kombinasi"] * 100).toFixed(2);
                pasaran_depandisckembar_5050kombinasi_field = (record[i]["depandisckembar_5050kombinasi"] * 100).toFixed(2);
                pasaran_limitglobal_5050kombinasi_field = record[i]["limitglobal_5050kombinasi"];
                pasaran_limittotal_5050kombinasi_field = record[i]["limittotal_5050kombinasi"];
                pasaran_minbet_kombinasi_field = record[i]["minbet_kombinasi"];
                pasaran_maxbet_kombinasi_field = record[i]["maxbet_kombinasi"];
                pasaran_win_kombinasi_field = record[i]["win_kombinasi"].toFixed(2);
                pasaran_disc_kombinasi_field = (record[i]["disc_kombinasi"] * 100).toFixed(2);
                pasaran_limitglobal_kombinasi_field = record[i]["limitglobal_kombinasi"];
                pasaran_limittotal_kombinasi_field = record[i]["limittotal_kombinasi"];
                
                pasaran_minbet_dasar_field = record[i]["minbet_dasar"];
                pasaran_maxbet_dasar_field = record[i]["maxbet_dasar"];
                pasaran_keibesar_dasar_field = (record[i]["keibesar_dasar"] * 100).toFixed(2);
                pasaran_keikecil_dasar_field = (record[i]["keikecil_dasar"] * 100).toFixed(2);
                pasaran_keigenap_dasar_field = (record[i]["keigenap_dasar"] * 100).toFixed(2);
                pasaran_keiganjil_dasar_field = (record[i]["keiganjil_dasar"] * 100).toFixed(2);
                pasaran_discbesar_dasar_field = (record[i]["discbesar_dasar"] * 100).toFixed(2);
                pasaran_disckecil_dasar_field = (record[i]["disckecil_dasar"] * 100).toFixed(2);
                pasaran_discgenap_dasar_field = (record[i]["discgenap_dasar"] * 100).toFixed(2);
                pasaran_discganjil_dasar_field = (record[i]["discganjil_dasar"] * 100).toFixed(2);
                pasaran_limitglobal_dasar_field = record[i]["limitglobal_dasar"];
                pasaran_limittotal_dasar_field = record[i]["limittotal_dasar"];
                pasaran_minbet_shio_field = record[i]["minbet_shio"];
                pasaran_maxbet_shio_field = record[i]["maxbet_shio"];
                pasaran_win_shio_field = (record[i]["win_shio"]).toFixed(2);
                pasaran_disc_shio_field = (record[i]["disc_shio"] * 100).toFixed(2);
                pasaran_shioyear_shio_field = record[i]["shioyear_shio"];
                pasaran_limitglobal_shio_field = record[i]["limitglobal_shio"];
                pasaran_limittotal_shio_field = record[i]["limittotal_shio"];
            }
        }
    }
    async function call_listpasaranonline() {
        listPasaranonline = [];
        const res = await fetch(path_api+"api/companypasaranonline", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                sData: "New",
                page: "COMPANY_HOME",
                company: idcompany,
                companypasaran_id: companypasaran_id,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listPasaranonline = [
                        ...listPasaranonline,
                        {
                            company_pasaranonline_id:record[i]["company_pasaranonline_id"],
                            company_pasaranonline_hari:record[i]["company_pasaranonline_hari"],
                        },
                    ];
                }
            }
        } 
    }
   
    async function handleSaveAdmin() {
        let flag = false;
        msg_error = "";
        admin_username_field_error = "";
        admin_password_field_error = "";
        admin_name_field_error = "";
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
            admin_username_field_error ="Format Username tidak sesuai pattern a-z atau 0-9"
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
            admin_name_field_error ="Format Name tidak sesuai pattern a-z atau A-Z atau 0-9 atau spasi"
        }
        if(admin_status_field == ""){
            flag = true
            admin_status_field_error ="Status is required"
        }
        if(flag == false){
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompanyadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataAdmin,
                    master: master,
                    company: idcompany,
                    admin_username: admin_username_field.toLowerCase,
                    admin_password: admin_password_field,
                    admin_name: admin_name_field,
                    admin_status: admin_status_field,
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
                        admin_username_field = "";
                        admin_password_field = "";
                        admin_name_field = "";
                        admin_status_field = "";
                        admin_username_field_error = "";
                        admin_password_field_error = "";
                        admin_name_field_error = "";
                        admin_status_field_error = "";
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
    async function saveupdatepasaran() {
        let flag = false;
        msg_error = "";
        if (pasaran_urlpasaran_field == "") {
            flag = true;
            msg_error += "The Pasaran URL is required<br>";
        }
        if (pasaran_pasarandiundi_field == "") {
            flag = true;
            msg_error += "The Pasaran diundi is required<br>";
        }
        if (pasaran_jamtutup_field == "") {
            flag = true;
            msg_error += "The Pasaran Jam Tutup is required<br>";
        }
        if (pasaran_jamjadwal_field == "") {
            flag = true;
            msg_error += "The Pasaran Jam Jadwal is required<br>";
        }
        if (pasaran_jamopen_field == "") {
            flag = true;
            msg_error += "The Pasaran Jam Open is required<br>";
        }
        if (pasaran_status_field == "") {
            flag = true;
            msg_error += "The Pasaran Status is required<br>";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompanyupdatepasaran", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    master: master,
                    company: idcompany,
                    companypasaran_id: companypasaran_id,
                    pasaran_diundi: pasaran_pasarandiundi_field,
                    pasaran_url: pasaran_urlpasaran_field,
                    pasaran_jamtutup: pasaran_jamtutup_field,
                    pasaran_jamjadwal: pasaran_jamjadwal_field,
                    pasaran_jamopen: pasaran_jamopen_field,
                    pasaran_statusactive: pasaran_status_field,
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
                call_listpasaran()
                call_listpasaranonline()
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    async function saveupdatepasaranline() {
        let flag = false;
        msg_error = "";
        if (pasaran_limitline4d_field == "") {
            flag = true;
            msg_error += "The Pasaran Limitline 4D is required\n";
        }
        if (pasaran_limitline3d_field == "") {
            flag = true;
            msg_error += "The Pasaran Limitline 3D is required\n";
        }
        if (pasaran_limitline3dd_field == "") {
            flag = true;
            msg_error += "The Pasaran Limitline 3DD is required\n";
        }
        if (pasaran_limitline2d_field == "") {
            flag = true;
            msg_error += "The Pasaran Limitline 2D is required\n";
        }
        if (pasaran_limitline2dd_field == "") {
            flag = true;
            msg_error += "The Pasaran Limitline 2DD is required\n";
        }
        if (pasaran_limitline2dt_field == "") {
            flag = true;
            msg_error += "The Pasaran Limitline 2DT is required\n";
        }
        if (pasaran_bbfs_field == "") {
            flag = true;
            msg_error += "The Pasaran BBFS is required\n";
        }
        if (flag == false) {
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompanyupdatepasaranline", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    master: master,
                    pasaran_id: pasaran_id_field,
                    company: idcompany,
                    companypasaran_id: companypasaran_id,
                    pasaran_limitline_4d: parseInt(pasaran_limitline4d_field),
                    pasaran_limitline_3d: parseInt(pasaran_limitline3d_field),
                    pasaran_limitline_3dd: parseInt(pasaran_limitline3dd_field),
                    pasaran_limitline_2d: parseInt(pasaran_limitline2d_field),
                    pasaran_limitline_2dd: parseInt(pasaran_limitline2dd_field),
                    pasaran_limitline_2dt: parseInt(pasaran_limitline2dt_field),
                    pasaran_bbfs: parseInt(pasaran_bbfs_field),
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
                call_listpasaranconf()
                call_listpasaranonline()
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true;
            }
        }
    }
    async function savePasaranOnline() {
        let flag = false;
        msg_error = "";
        if (select_pasaranonline == "" || select_pasaranonline == null) {
            flag = true;
            msg_error += "The Pasaran Online is required";
        }
        if (flag == false) {
            buttonLoading2_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/savecompanypasaranonline", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    master: master,
                    company: idcompany,
                    companypasaran_id: companypasaran_id,
                    pasaran_hari: select_pasaranonline,
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
                buttonLoading2_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                call_listpasaranconf()
                call_listpasaranonline()
            }
        } else {
            alert(msg);
        }
    }
    async function removeharionline(e) {
        const res = await fetch(path_api+"api/deletecompanypasaranonline", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                company: idcompany,
                companypasaran_id: parseInt(companypasaran_id),
                companypasaran_idoffline: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            call_listpasaranonline();
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
    const NewDataAdmin = (e,usernameadmincomp,nmadmincomp,stausadmincomp) => {
        sDataAdmin = e;
        modal_listadmin_width = "max-w-xl"
        isModal_Form_admin = true;
        if(e == "Edit"){
            admin_username_enable_field = false
            admin_username_field = usernameadmincomp;
            admin_password_field = "";
            admin_name_field = nmadmincomp;
            admin_status_field = stausadmincomp;
        }else{
            admin_username_enable_field = true
            admin_username_field = "";
            admin_password_field = "";
            admin_name_field = "";
            admin_status_field = "";
            admin_username_field_error = "";
            admin_password_field_error = "";
            admin_name_field_error = "";
            admin_status_field_error = "";
        }
    };
    
    const ChangeTabMenu = (e) => {
        switch(e){
            case "menu_listadmin":
                tab_listadmin = "bg-sky-600 text-white"
                tab_listpasaran = ""
                panel_listadmin = true
                panel_listpasaran = false
                break;
            case "menu_listpasaran":
                call_listpasaran();
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
    function clearFieldPasaran(){
        select_pasaranonline = "";
        companypasaran_id = "";
        pasaran_id_field = "";
        pasaran_nmpasarantogel_field = "";
        pasaran_urlpasaran_field = "";
        pasaran_pasarandiundi_field = "";
        pasaran_jamtutup_field = "";
        pasaran_jamjadwal_field = "";
        pasaran_jamopen_field = "";
        pasaran_status_field = "";

        pasaran_limitline4d_field = 0;
        pasaran_limitline3d_field = 0;
        pasaran_limitline3dd_field = 0;
        pasaran_limitline2d_field = 0;
        pasaran_limitline2dd_field = 0;
        pasaran_limitline2dt_field = 0;
        pasaran_bbfs_field = 0;
        pasaran_minbet_432d_field = 0;
        pasaran_maxbet4d_432d_field = 0;
        pasaran_maxbet3d_432d_field = 0;
        pasaran_maxbet3dd_432d_field = 0;
        pasaran_maxbet2d_432d_field = 0;
        pasaran_maxbet2dd_432d_field = 0;
        pasaran_maxbet2dt_432d_field = 0;
        pasaran_limitotal4d_432d_field = 0;
        pasaran_limitotal3d_432d_field = 0;
        pasaran_limitotal3dd_432d_field = 0;
        pasaran_limitotal2d_432d_field = 0;
        pasaran_limitotal2dd_432d_field = 0;
        pasaran_limitotal2dt_432d_field = 0;
        pasaran_limitglobal4d_432d_field = 0;
        pasaran_limitglobal3d_432d_field = 0;
        pasaran_limitglobal3dd_432d_field = 0;
        pasaran_limitglobal2d_432d_field = 0;
        pasaran_limitglobal2dd_432d_field = 0;
        pasaran_limitglobal2dt_432d_field = 0;
        pasaran_disc4d_432d_field = 0;
        pasaran_disc3d_432d_field = 0;
        pasaran_disc3dd_432d_field = 0;
        pasaran_disc2d_432d_field = 0;
        pasaran_disc2dd_432d_field = 0;
        pasaran_disc2dt_432d_field = 0;
        pasaran_win4d_432d_field = 0;
        pasaran_win3d_432d_field = 0;
        pasaran_win3dd_432d_field = 0;
        pasaran_win2d_432d_field = 0;
        pasaran_win2dd_432d_field = 0;
        pasaran_win2dt_432d_field = 0;
        pasaran_win4d_nodisc_432d_field = 0;
        pasaran_win3d_nodisc_432d_field = 0;
        pasaran_win3dd_nodisc_432d_field = 0;
        pasaran_win2d_nodisc_432d_field = 0;
        pasaran_win2dd_nodisc_432d_field = 0;
        pasaran_win2dt_nodisc_432d_field = 0;
        pasaran_win4d_bb_kena_432d_field = 0;
        pasaran_win3d_bb_kena_432d_field = 0;
        pasaran_win3dd_bb_kena_432d_field = 0;
        pasaran_win2d_bb_kena_432d_field = 0;
        pasaran_win2dd_bb_kena_432d_field = 0;
        pasaran_win2dt_bb_kena_432d_field = 0;
        pasaran_win4d_bb_432d_field = 0;
        pasaran_win3d_bb_432d_field = 0;
        pasaran_win3dd_bb_432d_field = 0;
        pasaran_win2d_bb_432d_field = 0;
        pasaran_win2dd_bb_432d_field = 0;
        pasaran_win2dt_bb_432d_field = 0;
        pasaran_minbet_cbebas_field = 0;
        pasaran_maxbet_cbebas_field = 0;
        pasaran_limitotal_cbebas_field = 0;
        pasaran_limitglobal_cbebas_field = 0;
        pasaran_win_cbebas_field = 0;
        pasaran_disc_cbebas_field = 0;
        pasaran_minbet_cmacau_field = 0;
        pasaran_maxbet_cmacau_field = 0;
        pasaran_limitotal_cmacau_field = 0;
        pasaran_limitglobal_cmacau_field = 0;
        pasaran_win2_cmacau_field = 0;
        pasaran_win3_cmacau_field = 0;
        pasaran_win4_cmacau_field = 0;
        pasaran_disc_cmacau_field = 0;
        pasaran_minbet_cnaga_field = 0;
        pasaran_maxbet_cnaga_field = 0;
        pasaran_win3_cnaga_field = 0;
        pasaran_win4_cnaga_field = 0;
        pasaran_disc_cnaga_field = 0;
        pasaran_limitglobal_cnaga_field = 0;
        pasaran_limittotal_cnaga_field = 0;
        pasaran_minbet_cjitu_field = 0;
        pasaran_maxbet_cjitu_field = 0;
        pasaran_winas_cjitu_field = 0;
        pasaran_winkop_cjitu_field = 0;
        pasaran_winkepala_cjitu_field = 0;
        pasaran_winekor_cjitu_field = 0;
        pasaran_desc_cjitu_field = 0;
        pasaran_limitglobal_cjitu_field = 0;
        pasaran_limittotal_cjitu_field = 0;
        pasaran_minbet_5050umum_field = 0;
        pasaran_maxbet_5050umum_field = 0;
        pasaran_keibesar_5050umum_field = 0;
        pasaran_keikecil_5050umum_field = 0;
        pasaran_keigenap_5050umum_field = 0;
        pasaran_keiganjil_5050umum_field = 0;
        pasaran_keitengah_5050umum_field = 0;
        pasaran_keitepi_5050umum_field = 0;
        pasaran_discbesar_5050umum_field = 0;
        pasaran_disckecil_5050umum_field = 0;
        pasaran_discgenap_5050umum_field = 0;
        pasaran_discganjil_5050umum_field = 0;
        pasaran_disctengah_5050umum_field = 0;
        pasaran_disctepi_5050umum_field = 0;
        pasaran_limitglobal_5050umum_field = 0;
        pasaran_limittotal_5050umum_field = 0;
        pasaran_minbet_5050special_field = 0;
        pasaran_maxbet_5050special_field = 0;
        pasaran_keiasganjil_5050special_field = 0;
        pasaran_keiasgenap_5050special_field = 0;
        pasaran_keiasbesar_5050special_field = 0;
        pasaran_keiaskecil_5050special_field = 0;
        pasaran_keikopganjil_5050special_field = 0;
        pasaran_keikopgenap_5050special_field = 0;
        pasaran_keikopbesar_5050special_field = 0;
        pasaran_keikopkecil_5050special_field = 0;
        pasaran_keikepalaganjil_5050special_field = 0;
        pasaran_keikepalagenap_5050special_field = 0;
        pasaran_keikepalabesar_5050special_field = 0;
        pasaran_keikepalakecil_5050special_field = 0;
        pasaran_keiekorganjil_5050special_field = 0;
        pasaran_keiekorgenap_5050special_field = 0;
        pasaran_keiekorbesar_5050special_field = 0;
        pasaran_keiekorkecil_5050special_field = 0;
        pasaran_discasganjil_5050special_field = 0;
        pasaran_discasgenap_5050special_field = 0;
        pasaran_discasbesar_5050special_field = 0;
        pasaran_discaskecil_5050special_field = 0;
        pasaran_disckopganjil_5050special_field = 0;
        pasaran_disckopgenap_5050special_field = 0;
        pasaran_disckopbesar_5050special_field = 0;
        pasaran_disckopkecil_5050special_field = 0;
        pasaran_disckepalaganjil_5050special_field = 0;
        pasaran_disckepalagenap_5050special_field = 0;
        pasaran_disckepalabesar_5050special_field = 0;
        pasaran_disckepalakecil_5050special_field = 0;
        pasaran_discekorganjil_5050special_field = 0;
        pasaran_discekorgenap_5050special_field = 0;
        pasaran_discekorbesar_5050special_field = 0;
        pasaran_discekorkecil_5050special_field = 0;
        pasaran_limitglobal_5050special_field = 0;
        pasaran_limittotal_5050special_field = 0;
        pasaran_minbet_5050kombinasi_field = 0;
        pasaran_maxbet_5050kombinasi_field = 0;
        pasaran_belakangkeimono_5050kombinasi_field = 0;
        pasaran_belakangkeistereo_5050kombinasi_field = 0;
        pasaran_belakangkeikembang_5050kombinasi_field = 0;
        pasaran_belakangkeikempis_5050kombinasi_field = 0;
        pasaran_belakangkeikembar_5050kombinasi_field = 0;
        pasaran_tengahkeimono_5050kombinasi_field = 0;
        pasaran_tengahkeistereo_5050kombinasi_field = 0;
        pasaran_tengahkeikembang_5050kombinasi_field = 0;
        pasaran_tengahkeikempis_5050kombinasi_field = 0;
        pasaran_tengahkeikembar_5050kombinasi_field = 0;
        pasaran_depankeimono_5050kombinasi_field = 0;
        pasaran_depankeistereo_5050kombinasi_field = 0;
        pasaran_depankeikembang_5050kombinasi_field = 0;
        pasaran_depankeikempis_5050kombinasi_field = 0;
        pasaran_depankeikembar_5050kombinasi_field = 0;
        pasaran_belakangdiscmono_5050kombinasi_field = 0;
        pasaran_belakangdiscstereo_5050kombinasi_field = 0;
        pasaran_belakangdisckembang_5050kombinasi_field = 0;
        pasaran_belakangdisckempis_5050kombinasi_field = 0;
        pasaran_belakangdisckembar_5050kombinasi_field = 0;
        pasaran_tengahdiscmono_5050kombinasi_field = 0;
        pasaran_tengahdiscstereo_5050kombinasi_field = 0;
        pasaran_tengahdisckembang_5050kombinasi_field = 0;
        pasaran_tengahdisckempis_5050kombinasi_field = 0;
        pasaran_tengahdisckembar_5050kombinasi_field = 0;
        pasaran_depandiscmono_5050kombinasi_field = 0;
        pasaran_depandiscstereo_5050kombinasi_field = 0;
        pasaran_depandisckembang_5050kombinasi_field = 0;
        pasaran_depandisckempis_5050kombinasi_field = 0;
        pasaran_depandisckembar_5050kombinasi_field = 0;
        pasaran_limitglobal_5050kombinasi_field = 0;
        pasaran_limittotal_5050kombinasi_field = 0;
        pasaran_minbet_kombinasi_field = 0;
        pasaran_maxbet_kombinasi_field = 0;
        pasaran_win_kombinasi_field = 0;
        pasaran_disc_kombinasi_field = 0;
        pasaran_limitglobal_kombinasi_field = 0;
        pasaran_limittotal_kombinasi_field = 0;
        pasaran_minbet_dasar_field = 0;
        pasaran_maxbet_dasar_field = 0;
        pasaran_keibesar_dasar_field = 0;
        pasaran_keikecil_dasar_field = 0;
        pasaran_keigenap_dasar_field = 0;
        pasaran_keiganjil_dasar_field = 0;
        pasaran_discbesar_dasar_field = 0;
        pasaran_disckecil_dasar_field = 0;
        pasaran_discgenap_dasar_field = 0;
        pasaran_discganjil_dasar_field = 0;
        pasaran_limitglobal_dasar_field = 0;
        pasaran_limittotal_dasar_field = 0;
        pasaran_minbet_shio_field = 0;
        pasaran_maxbet_shio_field = 0;
        pasaran_win_shio_field = 0;
        pasaran_disc_shio_field = 0;
        pasaran_shioyear_shio_field = "";
        pasaran_limitglobal_shio_field = 0;
        pasaran_limittotal_shio_field = 0;
    }
    const LoadingRunning = () => {
        loader_class = "inline-block"
        loader_msg = "Sending..."
    };
    const LoadingRunningFinish = (e) => {
        loader_msg =e.detail.temp_msg
        setTimeout(function () {
            loader_class = "hidden";
        }, 1000);
    };
    const call_notif = (e) => {
        msg_error = e.detail.temp_msg
        isModalNotif = true;
    };   
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
                    item.company_admin_username
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
    modal_popup_class="select-none w-11/12 {modal_width} overflow-hidden">
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
                            class="items-center {tab_listpasaran} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Pasaran</li>
                    </ul>
                    {#if panel_listadmin}
                        <h2 class="text-lg font-bold">List Admin</h2>
                        <div class="flex justify-between items-center w-full gap-1">
                            <input 
                                bind:value={searchListAdmin}
                                type="text" placeholder="Search by Username" class="input input-bordered w-full  rounded-md  focus:ring-0 focus:outline-none">
                            <button on:click={() => {
                                NewDataAdmin("New","","","");
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
                                                        NewDataAdmin("Edit",rec.company_admin_username,rec.company_admin_nama,rec.company_admin_status);
                                                    }}>
                                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                                        <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                                                    </svg>
                                                </td>
                                                <td class="text-xs text-center align-top">
                                                    <span class="{rec.company_admin_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.company_admin_status}</span>
                                                </td>
                                                <td class="text-xs text-left align-top">{rec.company_admin_typeadmin}</td>
                                                <td class="text-xs text-center align-top">{rec.company_admin_lastlogin}</td>
                                                <td class="text-xs text-center align-top">{rec.company_admin_lastipaddres}</td>
                                                <td class="text-xs text-left align-top">{rec.company_admin_username}</td>
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
                    input_maxlength_text="30"
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
                <select class="select select-bordered w-full rounded-md focus:ring-0 focus:outline-none" bind:value="{admin_status_field}">
                    <option disabled selected value="">--Pilih Status--</option>
                    <option value="ACTIVE">ACTIVE</option>
                    <option value="DEACTIVE">DEACTIVE</option>
                </select>
                {#if admin_status_field_error != ""}
                    <small class="text-pink-600 text-[11px]">{admin_status_field_error}</small>
                {/if}
            </div>
            <button
                on:click={() => {
                    handleSaveAdmin();
                }}  
                class="{buttonLoading_class} btn-block">Submit</button>
        </div>
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formpasaran" class="modal-toggle" bind:checked={isModal_Form_pasaran}>
<Modal_popup
    modal_popup_id="my-modal-formpasaran"
    modal_popup_title="Pasaran : {pasaran_nmpasarantogel_field}"
    modal_popup_class="select-none w-11/12 {modal_listpasaran_width} scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100">
    <slot:template slot="modalpopup_body">
        
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formconfpasaran" class="modal-toggle" bind:checked={isModal_Form_confpasaran}>
<Modal_popup
    modal_popup_id="my-modal-formconfpasaran"
    modal_popup_title="{pasaran_nmpasarantogel_field} - {permainan}"
    modal_popup_class="select-none w-11/12 {modal_confpasaran_width} scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100">
    <slot:template slot="modalpopup_body">
        
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />
