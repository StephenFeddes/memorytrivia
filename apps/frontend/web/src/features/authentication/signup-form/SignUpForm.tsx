import React from 'react';
import { useState } from 'react';
import { FormTextInput } from '../form-text-input/FormTextInput';
import styles from './SignUpForm.module.css';

export const SignUpForm = () => {
	const [username, setUsername] = useState('');
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');

	return (
		<form className={styles.form}>
			<FormTextInput
				placeholder="Username"
				handleOnChange={setUsername}
			/>
            <FormTextInput
                placeholder="Email"
                handleOnChange={setEmail}
            />
            <FormTextInput
                placeholder="Password"
                handleOnChange={setPassword}
            />
		</form>
	);
};
