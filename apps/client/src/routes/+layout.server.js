/** @type {import('./$types').LayoutServerLoad} */
export const load = (request_event) => {
    const accept_language = request_event.request?.headers?.get('accept-language') ?? "en"; 

    return  {
        language: processLocaleIdentifier(accept_language),
        is_virtual_mobile: determineIsMobile(request_event.request.headers),
    }
}

/**
 * @param {Headers} headers
 */
const determineIsMobile = headers => {
    let is_mobile = false;

    let user_agent = headers.get('user-agent') ?? '';
    user_agent = user_agent.toLowerCase();

    is_mobile = user_agent.includes('mobile');

    const sec_ch_ua_mobile = headers.get('sec-ch-ua-mobile') ?? '';

    if (sec_ch_ua_mobile) {
        is_mobile = sec_ch_ua_mobile === '?1';
    }

    return is_mobile;
}

const processLocaleIdentifier = identifier => {
    let selected_locale = 'en';

    let highest_quality = -0.1;

    const identifiers = identifier.split(', ');

    for (let h=0; h < identifiers.length; h++) {
        let [locale, q] = identifiers[h].split(';q=');

        q = q ?? 0;

        if (q > highest_quality) {
            highest_quality = q;
            selected_locale = locale;
        }
    }

    return selected_locale;
}