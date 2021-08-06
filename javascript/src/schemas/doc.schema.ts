import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type DocDocument = Doc & Document;

@Schema({collection: "docs"}) // The collection is called "docs"
export class Doc {
  @Prop()
  name: string;
}

export const DocSchema = SchemaFactory.createForClass(Doc);
