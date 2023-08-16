import { Select, SelectProps, Space } from 'antd';
import { getCategorySelectOptionByName } from 'container/NewWidget/RightContainer/alertFomatCategories';
import { useQueryBuilder } from 'hooks/queryBuilder/useQueryBuilder';

import { categoryToSupport } from './config';
import { DefaultLabel, selectStyles } from './styles';
import { IBuilderUnitsFilterProps } from './types';
import { filterOption } from './utils';

function BuilderUnitsFilter({
	onChange,
}: IBuilderUnitsFilterProps): JSX.Element {
	const { currentQuery, handleOnUnitsChange } = useQueryBuilder();

	const selectedValue = currentQuery?.unit;

	const allOptions = categoryToSupport.map((category) => ({
		label: category,
		options: getCategorySelectOptionByName(category),
	}));

	const onChangeHandler: SelectProps['onChange'] = (value): void => {
		if (onChange) {
			onChange(value);
		}

		handleOnUnitsChange(value);
	};

	return (
		<Space>
			<DefaultLabel>Y-axis unit</DefaultLabel>
			<Select
				style={selectStyles}
				onChange={onChangeHandler}
				value={selectedValue}
				options={allOptions}
				allowClear
				showSearch
				optionFilterProp="label"
				placeholder="Select unit"
				filterOption={filterOption}
			/>
		</Space>
	);
}

export { BuilderUnitsFilter };
