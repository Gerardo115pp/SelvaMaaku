
/** @type {import('./$types').PageServerLoad} */
export const load = ({ params, request }) => {
    const accept_language = request.headers.get('accept-language');

    return  {
        language: processLocaleIdentifier(accept_language),
    }
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