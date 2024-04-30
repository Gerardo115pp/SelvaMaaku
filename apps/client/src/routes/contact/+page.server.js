import { sendContactMail } from '@libs/libery_node_mailer/mailer';

/**
 * @type {import('./$types').Actions}
 */
export const actions = {
    submit: async http_request => {
        const form_data = await http_request.request.formData();

        const name = form_data.get('name');
        const email = form_data.get('email');
        const phone = form_data.get('phone');
        const message = form_data.get('message');

        let success = sendContactMail(name, email, phone, message);

        return {
            success
        }
    }
}