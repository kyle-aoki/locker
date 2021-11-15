# locker

Locker is a mini version of Hashicorp's Vault.

The idea is to simplify secrets management with an opinionated implementation of Vault. The idea begins with noticing a strong organization pattern emerge across companies that use Vault.

That pattern is essentially: the root folder containing repositories, where each folder is one-to-one with github repositories. For example, a github repository called `user-api` will have a folder `/user-api` in Vault, a github repository called `book-api` will have a folder `/book-api` in Vault, etc...

Then, within each project folder, there are environments. For example, a folder `user-api` will have environments `dev`, `stage`, and `prod`.

And finally, within each environment folder there are secrets. For example, `/user-api/dev` will have secret `/user-api/dev/POSTGRES_PASSWORD`, where each secret is mirrored across all environments. So, `POSTGRES_PASSWORD` will exist in `dev`, `stage`, and `prod`.

Vault is unopinionated in the sense that no one is forcing a company to organize their secrets in this manner. However, with such a strong pattern being seen across companies, it makes one wonder whether it is necessary to have such an open ended program like Vault. Why not have a more opinionated design, and in turn, have additional features which take advantage of that opinionated design?

Furthermore, Vault is quite complex, as it attempts to offer many features and integrations in order to capture the requirements of as many companies as possible. The consequence of this complexity is to cause Vault to be difficult to set up. Large companies will have an entire team dedicated to Vault, with `dev`, `stage`, and `prod` Vault environments to insure that company's Vault setup is correct.

Locker attempts to be as simple as possible. It accomplishes this by having only a server (Java Spring Boot) and a cli (golang).

The Spring Boot server uses Hibernate and JPA Repositories, which means server setup is as simple as attaching an SQL Database as environment variables.

Finally, another shortcoming of Vault is that no company needs such broad access to secrets. What I mean is: if a large company exists where `team A` is working on `Insurance Data`, and team B is working on `User Retention`, `team A` will *never* need access to `team B`'s secrets, and vice versa. Therefore it begs the question, why store both sets of secrets in the same Vault? The answer is that, because Vault is so complex, it would be a nuisance to set up one Vault for each logical division across a company.

Locker attempts to ameliorate this issue by providing a lightweight setup, so that one Locker per division can be provisioned. This reduces complexity and permissions bloat to a significant degree.
