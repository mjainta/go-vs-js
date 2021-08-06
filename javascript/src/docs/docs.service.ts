import { Model } from 'mongoose';
import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Doc, DocDocument } from '../schemas/doc.schema';

@Injectable()
export class DocsService {
  constructor(@InjectModel(Doc.name) private docModel: Model<DocDocument>) {}

  async findAll(): Promise<Doc[]> {
    return this.docModel.find().exec();
  }

  async findFiltered(name: string): Promise<Doc[]> {
    return this.docModel.find({
      name: {
        $regex: name
      }
    }).exec();
  }
}