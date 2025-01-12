# Определяем пути
base_dir="db/postgres/init"
output_file="build/schema.sql"

# Удаляем существующий файл, если он есть
if [ -f "$output_file" ]; then
    echo "Removing existing file: $output_file"
    rm "$output_file"
fi

# Создаем директорию, если она не существует
mkdir -p "$(dirname "$output_file")"
echo "Directory created or already exists: $(dirname "$output_file")"

# Создаем новый файл (файл будет создан с нуля)
touch "$output_file"
if [ -f "$output_file" ]; then
    echo "File created: $output_file"
else
    echo "Failed to create file: $output_file"
    exit 1
fi

find "$base_dir" -type f -name "tables.sql" -exec cat {} + > "$output_file"
echo "Content from sql files has been merged into $output_file"