<script>
    import ThemeButton from "@components/buttons/theme_button.svelte";
    import { site_language, supported_languages } from "@stores/site_content";
</script>

<form action="" id="smk-contact-form">    
    <label class="smk-cform-field">
        <span>
            {$site_language !== supported_languages.SPANISH ? "Name" : "Nombre"}
        </span>
        <div class="smk-cform-field-input-wrapper">
            <input 
                type="text"
                placeholder={$site_language !== supported_languages.SPANISH ? "Your name" : "Tu nombre"}
            >
        </div>
    </label>
    <label class="smk-cform-field">
        <span>
            <!-- Email -->
            {$site_language !== supported_languages.SPANISH ? "Email" : "Correo electrónico"}
        </span>
        <div class="smk-cform-field-input-wrapper">
            <svg viewBox="0 0 20 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M18.3337 2.99992C18.3337 2.08325 17.5837 1.33325 16.667 1.33325H3.33366C2.41699 1.33325 1.66699 2.08325 1.66699 2.99992M18.3337 2.99992V12.9999C18.3337 13.9166 17.5837 14.6666 16.667 14.6666H3.33366C2.41699 14.6666 1.66699 13.9166 1.66699 12.9999V2.99992M18.3337 2.99992L10.0003 8.83325L1.66699 2.99992" stroke="black" stroke-opacity="0.48" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>            
            <input 
                type="email" 
                placeholder={$site_language !== supported_languages.SPANISH ? "Your email" : "Tu correo electrónico"}
            >
        </div>
    </label>
    <label class="smk-cform-field">
        <span>
            <!-- Phone(width country code) -->
            {$site_language !== supported_languages.SPANISH ? "Phone(width country code)" : "Teléfono(con código de país)"}
        </span>
        <div class="smk-cform-field-input-wrapper">
            <svg viewBox="0 0 19 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M17.3332 14.0999V16.5999C17.3341 16.832 17.2866 17.0617 17.1936 17.2744C17.1006 17.487 16.9643 17.6779 16.7933 17.8348C16.6222 17.9917 16.4203 18.1112 16.2005 18.1855C15.9806 18.2599 15.7477 18.2875 15.5165 18.2666C12.9522 17.988 10.489 17.1117 8.32486 15.7083C6.31139 14.4288 4.60431 12.7217 3.32486 10.7083C1.91651 8.53426 1.04007 6.05908 0.76653 3.48325C0.745705 3.25281 0.773092 3.02055 0.846947 2.80127C0.920801 2.58199 1.03951 2.38049 1.1955 2.2096C1.3515 2.03871 1.54137 1.90218 1.75302 1.80869C1.96468 1.7152 2.19348 1.6668 2.42486 1.66658H4.92486C5.32928 1.6626 5.72136 1.80582 6.028 2.06953C6.33464 2.33324 6.53493 2.69946 6.59153 3.09992C6.69705 3.89997 6.89274 4.68552 7.17486 5.44158C7.28698 5.73985 7.31125 6.06401 7.24478 6.37565C7.17832 6.68729 7.02392 6.97334 6.79986 7.19992L5.74153 8.25825C6.92783 10.3445 8.65524 12.072 10.7415 13.2583L11.7999 12.1999C12.0264 11.9759 12.3125 11.8215 12.6241 11.755C12.9358 11.6885 13.2599 11.7128 13.5582 11.8249C14.3143 12.107 15.0998 12.3027 15.8999 12.4083C16.3047 12.4654 16.6744 12.6693 16.9386 12.9812C17.2029 13.2931 17.3433 13.6912 17.3332 14.0999Z" stroke="black" stroke-opacity="0.48" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
            <input 
                type="tel" 
                placeholder={$site_language !== supported_languages.SPANISH ? "Please include you'r country code" : "Por favor incluye tu código de país"}
            >
        </div>
    </label>
    <label class="smk-cform-field">
        <span>
            <!-- Message -->
            {$site_language !== supported_languages.SPANISH ? "Message" : "Mensaje"}
        </span>
        <div class="smk-cform-field-input-wrapper">
            <textarea></textarea>
        </div>
    </label>
    <ThemeButton 
        text={$site_language !== supported_languages.SPANISH ? "Send" : "Enviar"}
        is_button_one
    />
</form>

<style>
    #smk-contact-form {
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        row-gap: var(--spacing-3);
    }

    label.smk-cform-field {
        width: 100%;
        display: flex;
        flex-direction: column;
        row-gap: var(--spacing-1);
    }

    .smk-cform-field span {
        font-size: var(--font-size-h6);
    }

    .smk-cform-field-input-wrapper {
        width: 100%;
        display: flex;
        align-items: center;
        gap: var(--spacing-1);
        padding: var(--spacing-1) var(--spacing-2)  ;
        border-radius: var(--border-radius);
        background: var(--light-orange-light);
    }

    .smk-cform-field-input-wrapper:has(textarea) {
        height: 112px;
    }

    .smk-cform-field-input-wrapper svg {
        width: 19px;
    }

    .smk-cform-field input,
    .smk-cform-field textarea {
        width: 100%;
        padding: 0;
        border: 0;
        background: transparent;
    }

    .smk-cform-field textarea {
        resize: none;
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @container (width <= 800px) {
            .smk-cform-field span {
                font-size: var(--font-size-h3);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>