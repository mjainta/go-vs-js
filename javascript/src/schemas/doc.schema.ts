import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type DocDocument = Doc & Document;

@Schema({collection: "mycoll"}) // The collection is called "mycoll"
export class Doc {
  @Prop()
  name: string;
}

export const DocSchema = SchemaFactory.createForClass(Doc);
