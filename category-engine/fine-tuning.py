import pandas as pd
from datasets import load_dataset, Dataset
from transformers import BartTokenizer, BartForSequenceClassification, Trainer, TrainingArguments

df = pd.read_csv('fine-tuning.csv')

dataset = pd.DataFrame(df)
dataset = dataset.rename(columns={"title": "sentence", "category": "labels"})

print("Converted dataset")

dataset = Dataset.from_pandas(dataset)
train_test_split = dataset.train_test_split(test_size=0.2)

train_dataset = train_test_split['train']
test_dataset = train_test_split['test']

print("Splitted dataset into training & test datasets")

tokenizer = BartTokenizer.from_pretrained('facebook/bart-large-mnli')

def tokenize_function(examples):
    return tokenizer(examples['text'], padding='max_length', truncation=True)

train_dataset = train_dataset.map(tokenize_function, batched=True)
test_dataset = test_dataset.map(tokenize_function, batched=True)

print("Tokenized training & test datasets")

model = BartForSequenceClassification.from_pretrained('facebook/bart-large-mnli', num_labels=3)  # 3 categories: food, entertainment, grocery

training_args = TrainingArguments(
    output_dir='./results',          # output directory
    num_train_epochs=3,              # number of training epochs
    per_device_train_batch_size=8,   # batch size for training
    per_device_eval_batch_size=8,    # batch size for evaluation
    warmup_steps=500,                # number of warmup steps for learning rate scheduler
    weight_decay=0.01,               # strength of weight decay
    logging_dir='./logs',            # directory for storing logs
    logging_steps=10,
    eval_strategy="epoch",           # Evaluate the model every epoch
    save_strategy="epoch",           # Save the model every epoch
)

trainer = Trainer(
    model=model,                         # the model to be trained
    args=training_args,                  # training arguments
    train_dataset=train_dataset,         # training dataset
    eval_dataset=test_dataset,           # evaluation dataset
)

trainer.train()

print("Model trained")

model.save_pretrained('./fine_tuned_model')
tokenizer.save_pretrained('./fine_tuned_model')

print("Saved fine tuned model and tokens")

results = trainer.evaluate()
print("Evaluation results:", results)
