import React, { useState, useId } from 'react';
import styles from './FormTextInput.module.css';

interface Props {
    placeholder: string;
    type?: string;
    handleOnChange: (input: string) => void;
}

export const FormTextInput = ({
    placeholder,
    type = 'text',
    handleOnChange,
}: Props) => {
    const id = useId(); // Generates a unique ID
    const [showPassword, setShowPassword] = useState(type === 'password');

    return (
        <div className={styles.formTextInput}>
            <input
                id={id}
                type={showPassword ? 'password' : 'text'}
                placeholder={placeholder}
                onChange={(e) => handleOnChange(e.target.value)}
            />
            {type === 'password' && (
                <img
                    src={`icons/${showPassword ? 'eyehide' : 'eye'}.png`}
                    alt="Show password"
                    onClick={() => setShowPassword(!showPassword)}
                />
            )}
        </div>
    );
};
