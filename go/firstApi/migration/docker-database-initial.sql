create table personalities(
    id serial primary key,
    name varchar,
    description varchar
);

INSERT INTO personalities(name, description) VALUES
('Lorem Ipsum', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur non fermentum lorem. Donec quis ligula orci. Ut elementum erat vitae libero elementum iaculis. Ut aliquam, ante et interdum aliquet, odio sem ullamcorper turpis, sit amet semper ex erat quis augue. Nullam sodales euismod ipsum, sit amet pellentesque ipsum tristique et. Integer in justo arcu. Morbi maximus ipsum vitae consequat suscipit. Aenean lacus nisi, semper varius dictum a, lacinia in dui. Curabitur eleifend fermentum nisl id vehicula. Praesent mi urna, fermentum blandit tortor sed, dictum tempus est. Fusce id malesuada lorem. Etiam fermentum leo pharetra, sollicitudin nibh in, rhoncus felis.'),
('Sed porta erat', 'Sed porta erat et fermentum tempus. In vitae dui id neque pretium auctor. Suspendisse sollicitudin diam eu scelerisque rutrum. Proin euismod posuere porta. Donec maximus, justo efficitur dictum mattis, metus odio scelerisque lorem, vel egestas est eros vel mauris. Praesent pellentesque augue eu quam placerat tempus. Vestibulum quis imperdiet sapien, vitae luctus urna. Pellentesque ut efficitur diam, vel dapibus ipsum. Suspendisse condimentum nec nibh quis vulputate. Aliquam porttitor vitae libero a vehicula. Mauris et nibh id ante dignissim varius ac nec sapien. Morbi elementum fermentum dolor. Nam auctor, metus ac vulputate feugiat, elit est tempor felis, non tristique orci sapien quis nibh.');