ansible-playbook prod-deploy.yml \
--connection=local \
--tags=ec2-create \
--ask-vault-pass \
--extra-vars "state=stopped"