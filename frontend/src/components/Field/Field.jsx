import styles from './Field.module.scss'

function Field(props) {
    const {
        inputType = 'text',
        className= '',
        placeholder = '',
        onInput,
        id = '',
        autoComplete
    } = props

    return (
        <input 
            id={id}
            type={inputType} 
            placeholder={placeholder}
            className={`${styles.field} ${className}`}
            onInput={onInput}  
            autoComplete={autoComplete}  
        />
    )
}

export default Field