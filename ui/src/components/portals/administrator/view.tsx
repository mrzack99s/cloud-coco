import { FC, SetStateAction, useRef, useState } from 'react';
import { useForm, Controller } from 'react-hook-form';
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import { classNames } from 'primereact/utils';
import { useApiConnector } from '../../../utils/api_connector';

interface props {
    setDisplay: React.Dispatch<React.SetStateAction<any>>
}

export const View: FC<props> = ({ setDisplay }) => {
    const toast = useRef({} as any);
    const defaultValues = {
        name: '',
    }
    const [formData, setFormData] = useState(defaultValues);
    const [apiInstance] = useApiConnector()
    const { control, formState: { errors }, handleSubmit, reset } = useForm({ defaultValues });

    const getFormErrorMessage = (name: string) => {
        return errors.name && <small className="p-error">{errors.name.message}</small>
    };

    const onSubmit = (data: SetStateAction<{ name: string; }>) => {
        console.log(data)
        add(data.name)
    }

    const add = (name: string) => {
        apiInstance.services.createServices({
            name: name
        })
            .then(() => {
                toast.current.show({ severity: 'success', summary: 'Create a service success', life: 3000 });
                // refreshDataTable()
                setDisplay(false)
            })
            .catch((err) => {
                toast.current.show({ severity: 'error', summary: `Create a service not success`, detail: err.error.error, life: 3000 });
            })
    }

    return (
        <div>
            <Toast ref={toast} />
            <div className="flex justify-content-center">
                <div className="card w-10">
                    <form onSubmit={handleSubmit(onSubmit)} className="p-fluid">
                        <div className="field">
                            <span className="p-float-label">
                                <Controller name="name" control={control} rules={{ required: 'Name is required.' }} render={({ field, fieldState }) => (
                                    <InputText id={field.name} {...field} autoFocus className={classNames({ 'p-invalid': fieldState.invalid })} />
                                )} />
                                <label htmlFor="name" className={classNames({ 'p-error': errors.name })}>Name*</label>
                            </span>
                            {getFormErrorMessage('name')}
                        </div>
                        <Button type="submit" label="Submit" className="mt-2" />
                    </form>
                </div>
            </div>
        </div>
    );
}